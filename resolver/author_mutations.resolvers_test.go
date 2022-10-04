package resolver_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"

	"go-template/gqlmodels"
	"go-template/pkg/utl/convert"
	"go-template/resolver"
	"go-template/testutls"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestCreateAuthor(t *testing.T) {
     cases := []struct {
          name string
          req gqlmodels.AuthorCreateInput
          wantResp *gqlmodels.Author
          wantErr bool
     }{
          {
               name: "Fail on creating author",
               req: gqlmodels.AuthorCreateInput{},
               wantErr: true,
          },
          {
               name: "Success",
               req: gqlmodels.AuthorCreateInput{
                    FirstName: testutls.MockAuthor().FirstName.String,
                    LastName: testutls.MockAuthor().LastName.String,
                    Username: testutls.MockAuthor().Username.String,
               },
               wantResp: &gqlmodels.Author{
                    ID: fmt.Sprint(testutls.MockAuthor().ID),
                    FirstName: convert.NullDotStringToPointerString(testutls.MockAuthor().FirstName),
                    LastName: convert.NullDotStringToPointerString(testutls.MockAuthor().LastName),
                    Username: convert.NullDotStringToPointerString(testutls.MockAuthor().Username),
               },
               wantErr: false,
          },
     }
     resolver1 := resolver.Resolver{}
     for _, tt := range cases {
          t.Run(
               tt.name,
               func(t *testing.T) {
                    mock, db, _ := testutls.SetupEnvAndDB(
					t,
					testutls.Parameters{
						EnvFileLocation: "../.env.local",
					},
				)
				oldDB := boil.GetDB()
				defer func() {
					boil.SetDB(oldDB)
					db.Close()
				}()
				boil.SetDB(db)
                    if tt.name == "Fail on creating author" {
                         // insert new author
                         mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "authors"`)).
                              WithArgs().
                              WillReturnError(fmt.Errorf(""))
                    }
                    mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `authors` (`first_name`,`last_name`,`username`,`password`,`active`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?)")). //nolint
                         WithArgs(
                              testutls.MockAuthor().FirstName,
                              testutls.MockAuthor().LastName,
                              testutls.MockAuthor().Username,
                              "",
                              nil,
                              AnyTime{},
                              AnyTime{},
                              nil,
                         ).
                         WillReturnResult(sqlmock.NewResult(1, 1))

                    c := context.Background()
                    response, err := resolver1.Mutation().CreateAuthor(c, tt.req)
                    fmt.Println("RESPONSE: ", response)
                    fmt.Println("ERROR: ", err)
                    if tt.wantResp != nil {
                         assert.EqualValues(t, tt.wantResp, response)
                    }
                    assert.Equal(t, tt.wantErr, err != nil)
               },
          )
     }
}

func TestUpdateAuthor(t *testing.T){
     cases := []struct {
          name string
          req gqlmodels.AuthorUpdateInput
          wantResp *gqlmodels.Author
          wantErr bool
     }{
          {
               name: "Failed on updating author",
               req: gqlmodels.AuthorUpdateInput{},
               wantErr: true,
          }, 
          {
               name: "Success",
               req: gqlmodels.AuthorUpdateInput{
                    ID: "0",
                    FirstName: &testutls.MockAuthor().FirstName.String,
                    LastName: &testutls.MockAuthor().LastName.String,
                    Username: &testutls.MockAuthor().Username.String,
               },
               wantResp: &gqlmodels.Author{
                    ID: "0",
                    FirstName: &testutls.MockAuthor().FirstName.String,
                    LastName: &testutls.MockAuthor().LastName.String,
                    Username: &testutls.MockAuthor().Username.String,
               },
               wantErr: false,
          },
     }
     resolver1 := resolver.Resolver{}
     for _, tt := range cases {
          t.Run(
               tt.name,
               func (t *testing.T){
                    mock, db, _ := testutls.SetupEnvAndDB(
                         t, 
                         testutls.Parameters{
                              EnvFileLocation: "../.env.local",
                         })
                    oldDB := boil.GetDB()
                    defer func() {
                         db.Close()
                         boil.SetDB(oldDB)
                    }()
                    boil.SetDB(db)

                    if tt.wantErr {
                         mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "authors"`)).
                              WithArgs().
                              WillReturnError(fmt.Errorf(""))
                    }
                    rows := sqlmock.NewRows([]string{"first_name"}).
                                   AddRow(testutls.MockAuthor().FirstName)
                    mock.ExpectQuery(regexp.QuoteMeta("select * from `authors`")).
                         WithArgs(0).
                         WillReturnRows(rows)
                    mock.ExpectExec(regexp.QuoteMeta("UPDATE `authors` SET `first_name`=?,`last_name`=?,`username`=?,`password`=?,`active`=?,`updated_at`=?,`deleted_at`=? WHERE `id`=?")). //nolint
                                   WillReturnResult(sqlmock.NewResult(1, 1))

                    c := context.Background()
                    ctx := context.WithValue(c, testutls.AuthorKey, testutls.MockAuthor())
                    response, err := resolver1.Mutation().UpdateAuthor(ctx, &tt.req)
                    fmt.Println("RESPONSE: ", response)
                    fmt.Println("ERROR: ", err)
                    if tt.wantResp != nil && response != nil {
                         assert.Equal(t, tt.wantResp, response)
                    }
                    assert.Equal(t, tt.wantErr, err != nil)
               },
          )

     }
}

func TestDeleteAuthor(t *testing.T){
     cases := []struct {
          name string
          req *gqlmodels.AuthorDeleteInput
          wantResp *gqlmodels.AuthorDeletePayload
          wantErr bool
     }{
          {
               name: "Failed on deleting author",
               req: &gqlmodels.AuthorDeleteInput{},
               wantErr: true,
          },
          {
               name: "Success",
               req: &gqlmodels.AuthorDeleteInput{
                    ID: fmt.Sprint(testutls.MockAuthor().ID),
               },
               wantResp: &gqlmodels.AuthorDeletePayload{
                    ID: fmt.Sprint(testutls.MockAuthor().ID),
               },
               wantErr: false,
          },
     }
     resolver1 := resolver.Resolver{}
     query := regexp.QuoteMeta("select * from `authors` where `id`=?")
     for _, tt := range cases {
          t.Run(
               tt.name,
               func (t *testing.T){
                    mock, db, _ := testutls.SetupEnvAndDB(
                         t, 
                         testutls.Parameters{
                              EnvFileLocation: "../.env.local",
                         })
                    oldDB := boil.GetDB()
                    defer func(){
                         db.Close()
                         boil.SetDB(oldDB)
                    }()
                    boil.SetDB(db)

                    if tt.wantErr {
                         mock.ExpectQuery(query).WithArgs().WillReturnError(fmt.Errorf(""))
                    }
                    // get author by id
                    rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
                    mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)
                    // delete author
                    result := driver.Result(
                         driver.RowsAffected(
                              1,
                         ),
                    )
                    mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `authors` WHERE `id`=?")).
                         WillReturnResult(result)
                    
                    c := context.Background()
                    ctx := context.WithValue(c, testutls.AuthorKey, testutls.MockAuthor())
                    response, err := resolver1.Mutation().DeleteAuthor(ctx, tt.req)
                    fmt.Println("RESPONSE: ", response)
	               fmt.Println("ERROR: ", err)
                    if tt.wantResp != nil {
                         assert.Equal(t, tt.wantResp, response)
                    }
                    assert.Equal(t, tt.wantErr, err != nil)
               },
          )
     }
}