package daos_test

import (
	"context"
	"go-template/daos"
	"go-template/models"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestFindAuthorById(t *testing.T){
     cases := []struct {
          name string
          req int
          err error
     }{
          {
               name: "Passing an author_id",
               req: testutls.MockID,
               err: nil,
          },
     }
     mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
     oldDB := boil.GetDB()
     defer func(){
          boil.SetDB(oldDB)
          db.Close()
     }()
     boil.SetDB(oldDB)
     for _, tt := range cases {
          rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "username"}).
                         AddRow(testutls.MockAuthor().ID,
                              testutls.MockAuthor().FirstName,
                              testutls.MockAuthor().LastName,
                              testutls.MockAuthor().Username)
          mock.ExpectQuery(regexp.QuoteMeta("select * from `authors` where `id`=?")).
               WithArgs().
               WillReturnRows(rows)
          t.Run(tt.name, func(t *testing.T){
               _, err := daos.FindAuthorById(tt.req, context.Background())
               assert.Equal(t, err, tt.err)
          })
     }
}

func TestFindAllAuthorsWithCount(t *testing.T){
     cases := []struct {
          name string
          dbQueries []testutls.QueryData
          err error
     }{
          {
               name: "Success",
               err: nil,
               dbQueries: []testutls.QueryData{
                    {
                         Query: regexp.QuoteMeta("select * from `authors`;"),
                         DBResponse: sqlmock.NewRows([]string{"id", "first_name", "last_name", "username"}).
                                        AddRow(testutls.MockAuthor().ID,
                                             testutls.MockAuthor().FirstName,
                                             testutls.MockAuthor().LastName,
                                             testutls.MockAuthor().Username),
                    },
                    {
                         Query: regexp.QuoteMeta("select count(*) from `authors`;"),
                         DBResponse: sqlmock.NewRows([]string{"count"}).AddRow(testutls.MockCount),
                    },
               },
          },
     }
     mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
     oldDB := boil.GetDB()
     defer func(){
          boil.SetDB(oldDB)
          db.Close()
     }()
     boil.SetDB(oldDB)
     for _, tt := range cases {
          t.Run(tt.name, func(t *testing.T){
               for _, dbQuery := range tt.dbQueries {
                    mock.ExpectQuery(dbQuery.Query).
                         WithArgs().
                         WillReturnRows(dbQuery.DBResponse)

                    res, count, err := daos.FindAllAuthorsWithCount(context.Background())
                    if err != nil {
                         assert.Equal(t, true, err != nil)
                    } else {
                         assert.Equal(t, tt.err, err)
                         assert.Equal(t, count, testutls.MockCount)
                         assert.Equal(t, res[0].ID, testutls.MockAuthor().ID)
                         assert.Equal(t, res[0].FirstName, testutls.MockAuthor().FirstName)
                         assert.Equal(t, res[0].LastName, testutls.MockAuthor().LastName)
                         assert.Equal(t, res[0].Username, testutls.MockAuthor().Username)
                    }
               }
          })
     }
}

func TestCreateAuthor(t *testing.T){
     cases := []struct {
          name string
          req models.Author
          err error
     }{
          {
               name: "Success",
               req: models.Author{
                    ID: testutls.MockAuthor().ID,
                    FirstName: testutls.MockAuthor().FirstName,
                    LastName: testutls.MockAuthor().LastName,
                    Username: testutls.MockAuthor().Username,
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
     mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `authors`")).
          WithArgs().
          WillReturnResult(sqlmock.NewResult(1, 1))
     for _, tt := range cases {
          t.Run(tt.name, func(t *testing.T){
               _, err := daos.CreateAuthor(tt.req, context.Background())
               assert.Equal(t, err, tt.err)
          })
     }
}

func TestUpdateAuthor(t *testing.T){
     cases := []struct {
          name string
          req models.Author
          err error
     }{
          {
               name: "Success",
               req: models.Author{},
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
          mock.ExpectExec(regexp.QuoteMeta("UPDATE `authors` SET `first_name`=?,`last_name`=?,`username`=?,`password`=?,`active`=?,`updated_at`=?,`deleted_at`=? WHERE `id`=?")). //nolint
               WillReturnResult(sqlmock.NewResult(1, 1))
          t.Run(tt.name, func(t *testing.T){
               _, err := daos.UpdateAuthor(tt.req, context.Background())
               assert.Equal(t, tt.err, err)
          })
     }
}

func TestDeleteAuthor(t *testing.T){
     cases := []struct {
          name string
          req models.Author
          err error
     }{
          {
               name: "Success",
               req: models.Author{},
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
          mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `authors` WHERE `id`=?")).
               WithArgs().
               WillReturnResult(sqlmock.NewResult(1, 1))
          t.Run(tt.name, func(t *testing.T){
               _, err := daos.DeleteAuthor(tt.req, context.Background())
               assert.Equal(t, tt.err, err)
          })
     }
}