package resolver_test

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
	"go-template/resolver"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestLogin(t *testing.T){
     type args struct {
          Username string
          Password string
     }
     cases := []struct {
          name string
          req args
          err bool
          wantResp *gqlmodels.LoginResponse
     }{
          {
               name: "Fail on finding author by username",
               req: args{
                    Username: "wednesday",
                    Password: "123",
               },
               err: true,
          },
          {
               name: "Success",
               err: false,
               req: args{
                    Username: testutls.MockUsername,
                    Password: "adminuser",
               },
               wantResp: &gqlmodels.LoginResponse{
                    Token: "jwttokenstring",
                    RefreshToken: "refreshtoken",
               },
          },
     }
     mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../.env.local"})
     oldDB := boil.GetDB()
     defer func(){
          boil.SetDB(oldDB)
          db.Close()
     }()
     boil.SetDB(db)
     resolver1 := resolver.Resolver{}
     for _, tt := range cases {
          if tt.err {
               mock.ExpectQuery(regexp.QuoteMeta("SELECT `authors`.* FROM `authors` WHERE (username=?) LIMIT 1;")).
                    WithArgs().
                    WillReturnError(fmt.Errorf(""))
          }
          rows := sqlmock.NewRows([]string{"id", "password", "active"}).
                    AddRow(testutls.MockID, "$2a$10$dS5vK8hHmG5gzwV8f7TK5.WHviMBqmYQLYp30a3XvqhCW9Wvl2tOS", true)
          mock.ExpectQuery(regexp.QuoteMeta("SELECT `authors`.* FROM `authors` WHERE (username=?) LIMIT 1;")).
               WithArgs().
               WillReturnRows(rows)

          result := sqlmock.NewResult(1, 1)
          mock.ExpectExec(regexp.QuoteMeta("UPDATE `authors` ")).
               WillReturnResult(result)

          c := context.Background()
          response, err := resolver1.Mutation().Login(c, tt.req.Username, tt.req.Password)
          fmt.Println("RESP:", response)
          fmt.Println("ERR: ", err)
          if tt.wantResp != nil && response != nil {
               tt.wantResp.Token = response.Token
               tt.wantResp.RefreshToken = response.RefreshToken
               assert.Equal(t, tt.wantResp, response)
          }
          assert.Equal(t, tt.err, err != nil)
     }
}