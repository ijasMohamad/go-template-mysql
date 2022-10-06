package auth_test

import (
	"bytes"
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"go-template/gqlmodels"
	"go-template/internal/middleware/auth"
	"go-template/models"
	"go-template/resolver"
	"go-template/testutls"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var parseTokenMock func(token string) (*jwt.Token, error)

type tokenParseMock struct{
}

func (s tokenParseMock) ParseToken(token string) (*jwt.Token, error) {
	return parseTokenMock(token)
}

var operationHandlerMock func(ctx context.Context) graphql.ResponseHandler

func TestGraphQLMiddleware(t *testing.T) {
	cases := map[string]struct {
		wantStatus       int
		header           string
		signMethod       string
		err              string
		dbQueries        []testutls.QueryData
		operationHandler func(ctx context.Context) graphql.ResponseHandler
		tokenParser      func(token string) (*jwt.Token, error)
		whiteListedQuery bool
	}{
		"Success": {
			whiteListedQuery: false,
			header:           "Bearer 123",
			wantStatus:       http.StatusOK,
			err:              "",
			tokenParser: func(token string) (*jwt.Token, error) {
				return testutls.MockJWT(), nil
			},
			// tokenParser: func(token string) (*jwt.Token, error) {
			// 	return nil, nil
			// },
			operationHandler: func(ctx context.Context) graphql.ResponseHandler {
				author := ctx.Value(auth.AuthorCtxKey).(*models.Author)

				// Assertions
				assert.Equal(t, testutls.MockUsername, author.Username.String)
				assert.Equal(t, testutls.MockID, author.ID)
				assert.Equal(t, testutls.MockToken, author.Token.String)

				var handler = func(ctx context.Context) *graphql.Response {
					return &graphql.Response{
						Data: json.RawMessage([]byte("{}")),
					}
				}
				return handler
			},
			dbQueries: []testutls.QueryData{
				{
					Actions: &[]driver.Value{testutls.MockUsername},
					Query:   "SELECT `authors`.* FROM `authors` WHERE (username=?) LIMIT 1;",
					DBResponse: sqlmock.NewRows([]string{"id", "username", "token"}).
						AddRow(testutls.MockID, testutls.MockUsername, testutls.MockToken),
				},
			},
		},
		"Success__WhitelistedQuery": {
			whiteListedQuery: true,
			header:           "bearer 123",
			wantStatus:       http.StatusOK,
			err:              "",
			tokenParser: func(token string) (*jwt.Token, error) {
				// even without mocking the database or the token parser the middleware doesn't
				// throw an error since it skips all the checks and directly calls next
				return nil, nil
			},
			operationHandler: func(ctx context.Context) graphql.ResponseHandler {
				var handler = func(ctx context.Context) *graphql.Response {
					return &graphql.Response{
						Data: json.RawMessage([]byte(`{ "data": { "authors": { "id": 1 } } }`)),
					}
				}
				return handler
			},
			dbQueries: []testutls.QueryData{},
		},
		"Failure__NoAutherizationToken": {
			whiteListedQuery: false,
			header:           "",
			wantStatus:       http.StatusOK,
			err:              "Authorization header is missing",
			tokenParser: func(token string) (*jwt.Token, error) {
				return nil, nil
			},
			operationHandler: func(ctx context.Context) graphql.ResponseHandler {
				return nil
			},
			dbQueries: []testutls.QueryData{},
		},
		"Failure__InvalidAutherizationToken": {
			whiteListedQuery: false,
			header:           "bearer 123",
			wantStatus:       http.StatusOK,
			err:              "Invalid authorization token",
			tokenParser: func(token string) (*jwt.Token, error) {
				return nil, fmt.Errorf("token is invalid")
			},
			operationHandler: func(ctx context.Context) graphql.ResponseHandler {
				return nil
			},
			dbQueries: []testutls.QueryData{},
		},
		"Failure__NoAuthorWithThatUsername": {
			whiteListedQuery: false,
			header:           "bearer 123",
			wantStatus:       http.StatusOK,
			err:              "No author found for this username",
			tokenParser: func(token string) (*jwt.Token, error) {
				return testutls.MockJWT(), nil
			},
			operationHandler: func(ctx context.Context) graphql.ResponseHandler {
				return nil
			},
			dbQueries: []testutls.QueryData{},
		},
	}
	oldDB := boil.GetDB()
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../../../.env.local"})
	// boil.SetDB(db)

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			for _, dbQuery := range tt.dbQueries {
				mock.ExpectQuery(regexp.QuoteMeta(dbQuery.Query)).
					WithArgs(*dbQuery.Actions...).
					WillReturnRows(dbQuery.DBResponse)
			}
			requestQuery := testutls.MockQuery
			if tt.whiteListedQuery {
				requestQuery = testutls.MockWhiteListedQuery
			}
               fmt.Println("Name:", name)
			makeRequest(t, requestQuery, tt)
		})
	}
	boil.SetDB(oldDB)
	db.Close()
}

func makeRequest(t *testing.T, requestQuery string, tt struct {
	wantStatus       int
	header           string
	signMethod       string
	err              string
	dbQueries        []testutls.QueryData
	operationHandler func(ctx context.Context) graphql.ResponseHandler
	tokenParser      func(token string) (*jwt.Token, error)
	whiteListedQuery bool
}) {
	// mock token parser to handle the different cases for when the token is valid, invalid, empty
	parseTokenMock = tt.tokenParser

	// mock operation handler, and assert different conditions
	operationHandlerMock = tt.operationHandler

	tokenParser := tokenParseMock{}
	client := &http.Client{}
	observers := map[string]chan *gqlmodels.Author{}
	graphqlHandler := handler.New(gqlmodels.NewExecutableSchema(gqlmodels.Config{
		Resolvers: &resolver.Resolver{
			Observers:  observers,
		},
	}))
	graphqlHandler.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		res := auth.GraphQlMiddleware(ctx, tokenParser, operationHandlerMock)
		return res
	})
	graphqlHandler.AddTransport(transport.POST{})
	pathName := "/graphql"

	e := echo.New()
	e.POST(pathName, func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		graphqlHandler.ServeHTTP(res, req)
		return nil
	}, auth.GqlMiddleware())

	ts := httptest.NewServer(e)
	path := ts.URL + pathName
	defer ts.Close()

	req, _ := http.NewRequest(
		"POST",
		path,
		bytes.NewBuffer([]byte(requestQuery)),
	)

	if tt.wantStatus != 401 {
		req.Header.Set("authorization", tt.header)
	}
	req.Header.Set("Content-Type", "application/json")
     fmt.Println("Request:", req)

	res, err := client.Do(req)
	if err != nil {
          fmt.Println("Error after4:", err)
		t.Fatal("Cannot create http request")
	}
     fmt.Println("Client.Do")
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
          fmt.Println("Error after 4:", err)
		log.Fatal(err)
	}
     fmt.Println("bodyBytes")
	var jsonRes graphql.Response
	err = json.Unmarshal(bodyBytes, &jsonRes)

	if err != nil {
		log.Fatal(err)
	}
	for _, errorString := range jsonRes.Errors {
          // fmt.Println("RequestQuery:", requestQuery)
		fmt.Println("ErrorString:", errorString)
		fmt.Println("TT.Error:", tt.err)
		assert.Equal(t, tt.err, errorString.Message)
	}
	assert.Equal(t, tt.wantStatus, res.StatusCode)
}

func TestAuthorIdFromContext(t *testing.T){
     cases := map[string]struct {
          author *models.Author
          authorID int
     }{
          "Success": {
               author: &models.Author{ID: testutls.MockID},
               authorID: testutls.MockID,
          },
          "Failure": {
               author: nil,
               authorID: 0,
          },
     }
     for name, tt := range cases {
          t.Run(name, func(t *testing.T){
               authorID := auth.AuthorIdFromContext(context.WithValue(testutls.MockCtx{}, auth.AuthorCtxKey, tt.author))
               assert.Equal(t, tt.authorID, authorID)
          })
     }
}

func TestFromContext(t *testing.T){
     author := &models.Author{ID: testutls.MockID}
     a := auth.FromContext(context.WithValue(testutls.MockCtx{}, auth.AuthorCtxKey, author))
     assert.Equal(t, a, author)
     assert.Equal(t, author.ID, testutls.MockID)
}