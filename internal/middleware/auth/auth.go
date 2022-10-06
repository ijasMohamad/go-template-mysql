package auth

import (
	"context"
	"go-template/daos"
	"go-template/models"
	"go-template/pkg/utl/resultwrapper"
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/vektah/gqlparser/v2/ast"
)

type key string

var (
     authorization key = "Authorization"
)

type TokenParser interface {
     ParseToken(string) (*jwt.Token, error)
}
type CustomContext struct {
     echo.Context
     ctx context.Context
}

type ContextKey struct {
     name string
}
var AuthorCtxKey = &ContextKey{"author"}

func FromContext(ctx context.Context) *models.Author {
     author := ctx.Value(AuthorCtxKey).(*models.Author)
     return author
}

func AuthorIdFromContext(ctx context.Context) int {
     author := ctx.Value(AuthorCtxKey).(*models.Author)
     if author != nil {
          return author.ID
     }
     return 0
}

func GqlMiddleware () echo.MiddlewareFunc {
     return func(next echo.HandlerFunc) echo.HandlerFunc {
          return func(c echo.Context) error {
               ctx := context.WithValue(c.Request().Context(), authorization, c.Request().Header.Get(string(authorization)))
               c.SetRequest(c.Request().WithContext(ctx))
               cc := &CustomContext{c, ctx}
               return next(cc)
          }
     }
}

// WhiteListedOperations means graphql apis that can access without authentication
var WhiteListedOperations = map[string][]string {
     "query": {"__schema", "introspectionquery", "allArticles"},
     "mutation": {"login"},
}
func contains(s []string, e string) bool {
     for _, a := range s {
          if a == e {
               return true
          }
     }
     return false
}

func getAccessNeeds(operation *ast.OperationDefinition) (needsAuthAccess bool) {
     operationName := string(operation.Operation)
     for _, selectionSet := range operation.SelectionSet {
          selection := reflect.ValueOf(selectionSet).Elem()
          if !contains(WhiteListedOperations[operationName], selection.FieldByName("Name").Interface().(string)) {
               needsAuthAccess = true
          } 
     }
     return needsAuthAccess
}

func GraphQlMiddleware(ctx context.Context, tokenParser TokenParser, next graphql.OperationHandler) graphql.ResponseHandler {
 
     operation := graphql.GetOperationContext(ctx).Operation
     needsAuthAccess := getAccessNeeds(operation)
     if !needsAuthAccess {
          return next(ctx)
     }

     var tokenStr = ctx.Value(authorization).(string)
     if len(tokenStr) == 0 {
          return resultwrapper.HandleGraphQLError("Authorization header is missing")
     }

     token, err := tokenParser.ParseToken(tokenStr)
     
     if err != nil || !token.Valid {
          return resultwrapper.HandleGraphQLError("Invalid authorization token")
     }
     claims := token.Claims.(jwt.MapClaims)

     authorName := claims["u"].(string)

     author, err := daos.FindAuthorByUsername(authorName, ctx)
     if err != nil {
          return resultwrapper.HandleGraphQLError("No author found for this username")
     }
     ctx = context.WithValue(ctx, AuthorCtxKey, author)
     return next(ctx)
}