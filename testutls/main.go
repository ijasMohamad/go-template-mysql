package testutls

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"go-template/gqlmodels"
	"go-template/internal/config"
	"go-template/models"
	"go-template/pkg/utl/convert"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type key string

var AuthorKey key = "author"
var ArticleKey key = "article"

var MockID int = 1
var MockActive bool = false
var MockTitle string = "Title"
var MockCount = int(1)
var MockUsername = "username"
var MockToken = "token_string"
var MockJWTSecret = "1234567890123456789012345678901234567890123456789012345678901234567890"
var MockQuery = `{"query":"query allAuthors { allAuthors { authors { id } } }","variables":{}}"`
var MockWhiteListedQuery = `{"query":"query Schema { __schema { queryType { kind } } }","variables":{}}"`

func MockAuthor() *models.Author {
     return &models.Author {
          ID: MockID,
          FirstName: null.StringFrom("First"),
          LastName: null.StringFrom("Last"),
          Username: null.StringFrom("username"),
     }
}

func MockGqlAuthor() *gqlmodels.Author {
     return &gqlmodels.Author{
          ID: strconv.Itoa(MockID),
          FirstName: convert.StringToPointerString("First"),
          LastName: convert.StringToPointerString("Last"),
          Username: convert.StringToPointerString("username"),
          Active: &MockActive,
     }
}

func MockAuthors() []*models.Author {
     return []*models.Author {
          {
               FirstName: null.StringFrom("First"),
               LastName: null.StringFrom("Last"),
               Username: null.StringFrom("username"),
          },
     }
}

func MockArticle() *models.Article {
     return &models.Article{
          ID: MockID,
          Title: null.StringFrom("Title"),
          AuthorID: null.IntFrom(1),
          UpdatedAt: null.NewTime(time.Time{}, false),
          DeletedAt: null.NewTime(time.Time{}, false),
     }
}

func MockGqlArticle() *gqlmodels.Article {
     return &gqlmodels.Article{
          ID: strconv.Itoa(MockID),
          Title: convert.StringToPointerString("Title"),
     }
}

func MockArticles() []*models.Article {
     return []*models.Article {
         {
          Title: null.StringFrom("Title"),
         },
     }
} 

type Parameters struct {
     EnvFileLocation string `default:"../.env.local"`
}

func SetupEnv (envfile string) {
     err := godotenv.Load(envfile)
     if err != nil {
          fmt.Println("EnvFile:", envfile)
          fmt.Println("Error loading .env file", err)
     }
}

func SetupEnvAndDB (t *testing.T, parameters Parameters) (mock sqlmock.Sqlmock, db *sql.DB, err error) {
     SetupEnv(parameters.EnvFileLocation)
     db, mock, err = sqlmock.New()
     if err != nil {
          t.Fatalf("an error '%s' was not expected when opeining a stub database connection", err)
     }
     boil.SetDB(db)
     return mock, db, err
} 

func IsInTests () bool {
     for _, arg := range os.Args {
          if strings.HasPrefix(arg, "-test.paniconexit0") {
               return true
          }
     }
     return false
}

type QueryData struct {
     Actions *[]driver.Value
     Query string
     DBResponse *sqlmock.Rows
}

func MockConfig() *config.Configuration {
     return &config.Configuration{
          Server: &config.Server{
               Port: ":9000",
               Debug: true,
               ReadTimeout: 10,
               WriteTimeout: 5,
          },
          DB: &config.Database{
               LogQueries: true,
               Timeout: 5,
          },
          JWT: &config.JWT{
               MinSecretLength: 64,
               DurationMinutes: 1440,
               RefreshDuration: 3499200,
               MaxRefresh: 1440,
               SigningAlgorithm: "HS256",
          },
          App: &config.Application{
               MinPasswordStr: 1,
          },
     }
}

func MockJWT() *jwt.Token {
     return &jwt.Token{
          Raw: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIi" +
               "wibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
          Method: jwt.GetSigningMethod("HS256"),
          Claims: jwt.MapClaims{
               "u": MockUsername,
               "exp": "1.641189209e+09",
               "id": MockID,
          },
          Header: map[string]interface{}{
               "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
          },
          Valid: true,
     }
}