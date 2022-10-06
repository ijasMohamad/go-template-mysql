package daos_test

import (
	"context"
	"fmt"
	"go-template/daos"
	"go-template/models"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestFindArticleById(t *testing.T) {
     cases := []struct {
          name string
          req int
          err error
     }{
          {
               name: "Passing an article_id",
               req: 1,
               err: nil,
          },
     }
     for _, tt := range cases {
          err := godotenv.Load("../.env.local")
          if err != nil {
               fmt.Println("Error on loading env file")
          }
          db, mock, err := sqlmock.New()
          if err != nil {
               t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
          }
          // Inject mock instance into boil.
          oldDB := boil.GetDB()
          defer func(){
               db.Close()
               boil.SetDB(oldDB)
          }()
          boil.SetDB(db)

          rows := sqlmock.NewRows([]string {"id"}).AddRow(1)
          mock.ExpectQuery(regexp.QuoteMeta("select * from `articles` where `id`=?")).
               WithArgs().
               WillReturnRows(rows)

          t.Run(tt.name, func(t *testing.T){
               c := context.Background()
               _, err := daos.FindArticleById(1, c)
               assert.Equal(t, tt.err, err)
          })
     }
}

func TestFindAllArticlesWithCount(t *testing.T){
     oldDB := boil.GetDB()
     mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{
          EnvFileLocation: "../.env.local",
     })
     boil.SetDB(db)
     query := regexp.QuoteMeta("SELECT `articles`.* FROM `articles`;")
     cases := []struct {
          name string
          err error
          dbQueries []testutls.QueryData
     }{
          {
               name: "Failed to find all articles with count",
               err: fmt.Errorf("Error on finding articles with count"),
          },
          {
               name: "Success to find all articles with count",
               err: nil,
               dbQueries: []testutls.QueryData{
                    {
                         Query: query,
                         DBResponse: sqlmock.NewRows([]string{"id", "title"}).
                                        AddRow(testutls.MockID, testutls.MockTitle),
                    },
                    {
                         Query: regexp.QuoteMeta("SELECT COUNT(*) FROM `articles`;"),
                         DBResponse: sqlmock.NewRows([]string{"count"}).
                                        AddRow(testutls.MockCount),
                    },
               },
          },
     }
     for _, tt := range cases {
          if tt.err != nil {
               mock.ExpectQuery(query).WithArgs().WillReturnError(fmt.Errorf("Some error"))
          }

          for _, dbQuery := range tt.dbQueries {
               mock.ExpectQuery(dbQuery.Query).WithArgs().WillReturnRows(dbQuery.DBResponse)
               
               t.Run(tt.name, func(t *testing.T){
                    c := context.Background()
                    res, count, err := daos.FindAllArticlesWithCount(c)
                    if err != nil {
                         assert.Equal(t, true, err != nil)
                    } else {
                         assert.Equal(t, err, tt.err)
                         assert.Equal(t, testutls.MockCount, count)
                         assert.Equal(t, res[0].ID, testutls.MockID)
                         assert.Equal(t, res[0].Title, null.StringFrom(testutls.MockTitle))
                    }
               })
          } 
     }
     boil.SetDB(oldDB)
     db.Close()
}

func TestCreateArticle(t *testing.T){
     cases := []struct {
          name string
          req models.Article
          err error
     }{
          {
               name: "Success",
               req: models.Article{
                    ID: testutls.MockArticle().ID,
                    Title: testutls.MockArticle().Title,
                    AuthorID: testutls.MockArticle().AuthorID,
               },
               err: nil,
          },
     }
     mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
     oldDB := boil.GetDB()
     defer func(){
          boil.SetDB(oldDB)
          db.Close()
     }()
     boil.SetDB(db)

     for _, tt := range cases {
          t.Run(tt.name, func(t *testing.T){
               mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `articles`")).
                    WithArgs().
                    WillReturnResult(sqlmock.NewResult(1, 1))
               _, err := daos.CreateArticle(tt.req, context.Background())
               if err != nil {
                    assert.Equal(t, true, tt.err != nil)
               }
               assert.Equal(t, err, tt.err)
          })
     }
}

func TestUpdateArticle(t *testing.T){
     cases := []struct {
          name string
          req models.Article
          err error
     }{
          {
               name: "Success",
               req: models.Article{},
               err: nil,
          },
     }
     mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
     oldDB := boil.GetDB()
     defer func(){
          boil.SetDB(oldDB)
          db.Close()
     }()
     boil.SetDB(db)
     for _, tt := range cases {
          mock.ExpectExec(regexp.QuoteMeta("UPDATE `articles` SET `title`=?,`author_id`=?,`updated_at`=?,`deleted_at`=? WHERE `id`=?")). //nolint
                         WillReturnResult(sqlmock.NewResult(1, 1))

          t.Run(tt.name, func(t *testing.T){
               _, err := daos.UpdateArticle(tt.req, context.Background())
               assert.Equal(t, err, tt.err)
          })
     }
}

func TestDeleteArticle(t *testing.T){
     cases := []struct {
          name string
          req models.Article
          err error
     }{
          {
               name: "Success",
               req: models.Article{},
               err: nil,
          },
     }
     mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
     oldDB := boil.GetDB()
     defer func(){
          boil.SetDB(oldDB)
          db.Close()
     }()
     boil.SetDB(db)
     mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `articles` WHERE `id`=?")).
          WillReturnResult(sqlmock.NewResult(1, 1))

     for _, tt := range cases {
          t.Run(tt.name, func(t *testing.T){
               _, err := daos.DeleteArticle(tt.req, context.Background())
               assert.Equal(t, err, tt.err)
          })
     }
}
