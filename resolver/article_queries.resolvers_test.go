package resolver_test

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/convert"
	"go-template/resolver"
	"go-template/testutls"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestAllArticles(t *testing.T) {
     cases := []struct {
          name string
          wantResp []*models.Article
          wantErr bool
     }{
          {
               name: "Success",
               wantErr: false,
               wantResp: testutls.MockArticles(),
          },
     }
     resolver1 := resolver.Resolver{}
     for _, tt := range cases {
          t.Run(
               tt.name,
               func (t *testing.T) {
                    err := godotenv.Load("../.env.local")
                    if err != nil {
                         fmt.Println("Error loading .env file")
                    }
                    db, mock, err := sqlmock.New()
                    if err != nil {
                         t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
                    }
                    // Inject mock instance into boil
                    oldDB := boil.GetDB()
                    defer func () {
                         db.Close()
                         boil.SetDB(oldDB)
                    }()
                    boil.SetDB(db)
                    // Get article by id
                    rows := sqlmock.NewRows([]string{"id", "title"}).
                              AddRow(testutls.MockID, "Title")
                    mock.ExpectQuery(regexp.QuoteMeta("SELECT `articles`.* FROM `articles`;")).
                         WithArgs().
                         WillReturnRows(rows)

                    rowCount := sqlmock.NewRows([]string{"count"}).
                                   AddRow(1)
                    mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM `articles`;")).
                         WithArgs().
                         WillReturnRows(rowCount)

                    c := context.Background()
                    ctx := context.WithValue(c, testutls.ArticleKey, testutls.MockArticle())
                    response, err := resolver1.Query().AllArticles(ctx)
                    if response != nil && tt.wantResp != nil {
                         assert.Equal(t, len(tt.wantResp), len(response.Articles))
                    }
                    assert.Equal(t, tt.wantErr, err != nil)
               },
          )

     } 
}

func TestArticle(t *testing.T) {
     cases := []struct {
          name string
          req int
          wantResp *gqlmodels.Article
          wantErr bool
     }{
          {
               name: "Success",
               wantErr: false,
               req: 1,
               wantResp: &gqlmodels.Article{
                    ID: strconv.Itoa(testutls.MockArticle().ID),
                    Title: convert.NullDotStringToPointerString(testutls.MockArticle().Title),
               },
          },
     }
     resolver1 := resolver.Resolver{}
     for _, tt := range cases {
          t.Run(
               tt.name,
               func(t *testing.T) {
                    err := godotenv.Load("../.env.local")
                    if err != nil {
                         fmt.Println("Error loading .env file")
                    }
                    db, mock, err := sqlmock.New()
                    if err != nil {
                         t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
                    }
                    // Inject mock instance into boil
                    oldDB := boil.GetDB()
                    defer func() {
                         db.Close()
                         boil.SetDB(oldDB)
                    }()
                    boil.SetDB(db)
                    // Get article by id
                    rows := sqlmock.NewRows([]string{"id", "title"}).
                              AddRow(1, "Title")
                    mock.ExpectQuery(regexp.QuoteMeta("select * from `articles` where `id`=?")).
                         WithArgs(1).
                         WillReturnRows(rows)

                    c := context.Background()
                    ctx := context.WithValue(c, testutls.ArticleKey, testutls.MockArticle())
                    response, err := resolver1.Query().Article(ctx, 1)
                    fmt.Println("Response: ", response)
                    fmt.Println("Error :", err)

                    if tt.wantResp != nil && response != nil {
                         assert.Equal(t, tt.wantResp, response)
                    }
                    assert.Equal(t, tt.wantErr, err != nil)
               },
          )
     }
}