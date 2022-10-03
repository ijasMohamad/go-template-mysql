package testutls

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/convert"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
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

func MockAuthor() *models.Author {
     return &models.Author {
          ID: MockID,
          FirstName: null.StringFrom("First"),
          LastName: null.StringFrom("Last"),
          Username: null.StringFrom("username"),
          // Active: null.BoolFrom(false),
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
          fmt.Print("Error loading .env file")
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