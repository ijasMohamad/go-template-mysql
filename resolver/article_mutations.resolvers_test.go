package resolver_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"go-template/gqlmodels"
	"go-template/pkg/utl/convert"
	"go-template/resolver"
	"go-template/testutls"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
func TestCreateArticle(t *testing.T) {
	cases := []struct {
		name     string
		req      gqlmodels.ArticleCreateInput
		wantResp *gqlmodels.Article
		wantErr  bool
	}{
		{
		     name: "Fail on create article",
		     req: gqlmodels.ArticleCreateInput{},
		     wantErr: true,
		},
		{
			name: "Success",
			req: gqlmodels.ArticleCreateInput{
				Title:    testutls.MockArticle().Title.String,
				AuthorID: strconv.Itoa(testutls.MockArticle().AuthorID.Int),
			},
			wantResp: &gqlmodels.Article{
				ID:    fmt.Sprint(testutls.MockArticle().ID),
				Title: convert.NullDotStringToPointerString(testutls.MockArticle().Title),
			},
			wantErr: false,
		},
	}
	resolver1 := resolver.Resolver{}
	for _, tt := range cases {
		t.Run(
			tt.name,
			func(t *testing.T) {
				mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../.env.local"})
				oldDB := boil.GetDB()
				defer func() {
					db.Close()
					boil.SetDB(oldDB)
				}()
				boil.SetDB(db)
				if tt.name == "Fail on create article" {
					// insert new article
					mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "articles"`)).
						WithArgs().
						WillReturnError(fmt.Errorf(""))
				}

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `articles` (`title`,`author_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?)")). //nolint
                              WithArgs(
						testutls.MockArticle().Title,
						testutls.MockArticle().AuthorID,
						AnyTime{},
						AnyTime{},
						nil,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))

				c := context.Background()
				response, err := resolver1.Mutation().CreateArticle(c, tt.req)
				fmt.Println("Response: ", response)
				fmt.Println("Error: ", err)
				if tt.wantResp != nil {
					assert.EqualValues(t, tt.wantResp, response)
				}
				assert.Equal(t, tt.wantErr, err != nil)
			},
		)
	}
}

func TestUpdateArticle(t *testing.T) {
	cases := []struct {
		name     string
		req      gqlmodels.ArticleUpdateInput
		wantResp *gqlmodels.Article
		wantErr  bool
	}{
		{
			name:    "Fail on finding article",
			req:     gqlmodels.ArticleUpdateInput{},
			wantErr: true,
		},
		{
			name: "Success",
			req: gqlmodels.ArticleUpdateInput{
                    ID: "0",
				Title: convert.NullDotStringToPointerString(testutls.MockArticle().Title),
			},
			wantErr: false,
			wantResp: &gqlmodels.Article{
				ID:    "0",
				Title: convert.NullDotStringToPointerString(testutls.MockArticle().Title),
			},
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
					db.Close()
					boil.SetDB(oldDB)
				}()
				boil.SetDB(db)

				if tt.wantErr {
					mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "articles"`)).
						WithArgs().
						WillReturnError(fmt.Errorf(""))
				}
				rows := sqlmock.NewRows([]string{"title"}).
					AddRow(testutls.MockArticle().Title)
				mock.ExpectQuery(regexp.QuoteMeta("select * from `articles`")).
					WithArgs().
					WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `articles` SET `title`=?,`author_id`=?,`updated_at`=?,`deleted_at`=? WHERE `id`=?")). //nolint
                         WithArgs().
					WillReturnResult(sqlmock.NewResult(1, 1))

				c := context.Background()
				ctx := context.WithValue(c, testutls.ArticleKey, testutls.MockArticle())
				response, err := resolver1.Mutation().UpdateArticle(ctx, tt.req)
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

func TestDeleteArticle(t *testing.T) {
	cases := []struct {
		name     string
		req      *gqlmodels.ArticleDeleteInput
		wantResp *gqlmodels.ArticleDeletePayload
		wantErr  bool
	}{
		{
			name:    "Failed on deleting article",
			req:     &gqlmodels.ArticleDeleteInput{},
			wantErr: true,
		},
		{
			name: "Success",
			req: &gqlmodels.ArticleDeleteInput{
				ID: fmt.Sprint(testutls.MockAuthor().ID),
			},
			wantResp: &gqlmodels.ArticleDeletePayload{
				ID: fmt.Sprint(testutls.MockAuthor().ID),
			},
			wantErr: false,
		},
	}
	resolver1 := resolver.Resolver{}
	query := regexp.QuoteMeta("select * from `articles` where `id`=?")
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
					db.Close()
					boil.SetDB(oldDB)
				}()
				boil.SetDB(db)

				if tt.wantErr {
					mock.ExpectQuery(query).WithArgs().WillReturnError(fmt.Errorf(""))
				}
				// get article by id
				rows := mock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)
				// delete article by id
				result := driver.Result(
					driver.RowsAffected(
						1,
					),
				)
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `articles` WHERE `id`=?")).
					WithArgs(1).
					WillReturnResult(result)

				c := context.Background()
				ctx := context.WithValue(c, testutls.ArticleKey, testutls.MockGqlArticle())
				response, err := resolver1.Mutation().DeleteArticle(ctx, tt.req)
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
