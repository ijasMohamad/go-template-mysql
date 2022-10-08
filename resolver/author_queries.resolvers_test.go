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

func TestAuthor(t *testing.T) {
     cases := []struct {
          name string
          req int
          wantResp *gqlmodels.Author
          wantErr bool
     }{
          {
               name: "Success",
               req: 1,
               wantErr: false,
               wantResp: &gqlmodels.Author{
                    ID: strconv.Itoa(testutls.MockAuthor().ID),
                    FirstName: convert.NullDotStringToPointerString(testutls.MockAuthor().FirstName),
                    LastName: convert.NullDotStringToPointerString(testutls.MockAuthor().LastName),
                    Username: convert.NullDotStringToPointerString(testutls.MockAuthor().Username),
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
                         t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
                    }
                    // Inject mock instance into boil
                    oldDB := boil.GetDB()
                    defer func() {
                         db.Close()
                         boil.SetDB(oldDB)
                    }()
                    boil.SetDB(db)
                    
                    // get author by id
                    rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "active"}).
                                   AddRow(1, "First", "Last", "username", nil)
                    mock.ExpectQuery(regexp.QuoteMeta("select * from `authors` where `id`=?")).
                         WithArgs(1).
                         WillReturnRows(rows)

                    c := context.Background()
                    ctx := context.WithValue(c, testutls.AuthorKey, testutls.MockAuthor())
                    response, err := resolver1.Query().Author(ctx, 1)
                    fmt.Println("response: ", response)
                    fmt.Printf("err: %v\n\n", err)

                    if tt.wantResp != nil && response != nil {
                         assert.EqualValues(t, tt.wantResp, response)
                    }
                    assert.Equal(t, tt.wantErr, err != nil)
               },
          )
     }
}

func TestAllAuthors(t *testing.T) {
	cases := []struct {
		name     string
		wantResp []*models.Author
		wantErr  bool
	}{
		{
			name:     "Success",
			wantErr:  false,
			wantResp: testutls.MockAuthors(),
		},
	}
	resolver1 := resolver.Resolver{}
     for _, tt := range cases {
          t.Run(
               tt.name,
               func(t *testing.T) {
                    err := godotenv.Load("../.env.local")
                    if err != nil {
                         fmt.Print("Error loading .env file")
                    }
                    db, mock, err := sqlmock.New()
                    if err != nil {
                         t.Fatalf("An error '%s' was not expected when opening a stub database connection ", err)
                    }
                    // Inject mock instance into boil
                    oldDB := boil.GetDB()
                    defer func() {
                         db.Close()
                         boil.SetDB(oldDB)
                    }()
                    boil.SetDB(db)

                    // get author by id
                    rows := sqlmock.
                              NewRows(
                                   []string{"id", "first_name", "last_name", "username"},
                              ).
                              AddRow(testutls.MockID, "First", "Last", "username")
                              Query := regexp.QuoteMeta("SELECT `authors`.* FROM `authors`;")
                              fmt.Println("Query all: ", Query)

                    mock.ExpectQuery(Query).
                              WithArgs().
                              WillReturnRows(rows)
                    
                    rowCount := sqlmock.NewRows([]string{"count"}).AddRow(1)

                    mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM `authors`;")).
                                   WithArgs().
                                   WillReturnRows(rowCount)

                    c := context.Background()
                    ctx := context.WithValue(c, testutls.AuthorKey, testutls.MockAuthor())
                    
                    response, err := resolver1.Query().AllAuthors(ctx)
                    fmt.Println("response: ", response)

                    if tt.wantResp != nil && response != nil {
                         assert.Equal(t, len(tt.wantResp), len(response.Authors))
                    }
                    assert.Equal(t, tt.wantErr, err != nil )
               },
          )
     }
}
