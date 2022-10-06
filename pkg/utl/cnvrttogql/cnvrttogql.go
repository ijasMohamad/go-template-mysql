package cnvrttogql

import (
	"context"
	"strconv"

	"go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/convert"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func AuthorsToGraphQLAuthors(a models.AuthorSlice) []*gqlmodels.Author {
	var r []*gqlmodels.Author
	for _, v := range a {
		r = append(r, AuthorToGraphQLAuthor(v))
	}
	return r
}

func AuthorToGraphQLAuthor(a *models.Author) *gqlmodels.Author {
	if a == nil {
		return nil
	}
	var articles models.ArticleSlice
	a.L.LoadArticles(context.Background(), boil.GetContextDB(), true, a, nil) //nolint:errCheck
	if a.R != nil {
		articles = a.R.Articles
	}

	return &gqlmodels.Author{
		ID:        strconv.Itoa(a.ID),
		FirstName: convert.NullDotStringToPointerString(a.FirstName),
		LastName:  convert.NullDotStringToPointerString(a.LastName),
		Username:  convert.NullDotStringToPointerString(a.Username),
		Active:    convert.NullDotBoolToPointerBool(a.Active),
		Articles:  ArticlesToGraphQLArticles(articles),
	}
}

func ArticlesToGraphQLArticles(a models.ArticleSlice) []*gqlmodels.Article {
	var r []*gqlmodels.Article
	for _, v := range a {
		r = append(r, ArticleToGraphQLArticle(v))
	}
	return r
}

func ArticleToGraphQLArticle(a *models.Article) *gqlmodels.Article {
     if a == nil {
		return nil
	}
	return &gqlmodels.Article{
		ID:    strconv.Itoa(a.ID),
		Title: &a.Title.String,
	}
}
