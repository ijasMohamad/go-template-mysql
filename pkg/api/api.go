// Package api Go-Template
package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go-template/gqlmodels"
	graphql "go-template/gqlmodels"
	"go-template/internal/config"
	"go-template/internal/jwt"
	"go-template/internal/middleware/auth"
	"go-template/internal/mysql"
	"go-template/internal/server"
	"go-template/models"
	"go-template/resolver"

	graphql2 "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // here
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Start starts the API service
func Start(cfg *config.Configuration) (*echo.Echo, error) {
	fmt.Printf("connection started")

	db, err := mysql.Connect()
	if err != nil {
		return nil, err
	}

	boil.SetDB(db)

	jwt, err := jwt.New(
		cfg.JWT.SigningAlgorithm,
		os.Getenv("JWT_SECRET"),
		cfg.JWT.DurationMinutes,
		cfg.JWT.MinSecretLength,
	)
	if err != nil {
		return nil, err
	}

	e := server.New()

	gqlMiddleware := auth.GqlMiddleware()

	graphQLPathname := "/graphql"
	playgroundHandler := playground.Handler("GraphQL playground", graphQLPathname)

	observers := map[string]chan *graphql.Author{}
	observer2 := map[string]chan *graphql.Article{}

	c := graphql.Config{
		Resolvers: &resolver.Resolver{
			Observers:  observers,
			Observers2: observer2,
		},
	}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql2.Resolver, role gqlmodels.Role) (interface{}, error) { //nolint
		author := ctx.Value(auth.AuthorCtxKey).(*models.Author)
		// Checking the logged user is admin or not
		if role != graphql.Role(author.Role.String) {
			return nil, fmt.Errorf("Access denied")
		}
		return next(ctx)
	}

	graphqlHandler := handler.New(graphql.NewExecutableSchema(c))

	if os.Getenv("ENVIRONMENT_NAME") == "local" {
		boil.DebugMode = true
	}

	// Graphql apis
	graphqlHandler.AroundOperations(func(ctx context.Context, next graphql2.OperationHandler) graphql2.ResponseHandler {
		return auth.GraphQlMiddleware(ctx, jwt, next)
	})

	e.POST(graphQLPathname, func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		graphqlHandler.ServeHTTP(res, req)
		return nil
	}, gqlMiddleware)

	e.GET(graphQLPathname, func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		graphqlHandler.ServeHTTP(res, req)
		return nil
	}, gqlMiddleware)

	graphqlHandler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	graphqlHandler.AddTransport(transport.Options{})
	graphqlHandler.AddTransport(transport.GET{})
	graphqlHandler.AddTransport(transport.POST{})
	graphqlHandler.AddTransport(transport.MultipartForm{})

	graphqlHandler.SetQueryCache(lru.New(1000))

	graphqlHandler.Use(extension.Introspection{})
	graphqlHandler.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	// graphql playground
	e.GET("/playground", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		playgroundHandler.ServeHTTP(res, req)
		return nil
	})
	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})
	return e, nil
}
