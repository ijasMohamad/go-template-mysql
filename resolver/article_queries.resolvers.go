package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/cnvrttogql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, id int) (*gqlmodels.Article, error) {
	var contextExecutor = boil.GetContextDB()

	// article, err := models.FindArticle(ctx, contextExecutor, id)
	// if err != nil {
	// 	return nil, err
	// }
	// author, err := article.Author().One(ctx, contextExecutor)
	// if err != nil {
	// 	return nil, err
	// }
	// return cnvrttogql.ArticleToGraphQLArticle(article), nil

	article, err := models.FindArticle(ctx, contextExecutor, id)
	if err != nil {
		return nil, err
	}
	// author, err := models.FindAuthor(ctx, contextExecutor, article.AuthorID.Int)
	// if err != nil {
	// 	return nil, err
	// }
	// art := cnvrttogql.ArticleToGraphQLArticle(article)
	// art.Author = cnvrttogql.AuthorToGraphQLAuthor(author)
	// return art, nil

	return cnvrttogql.ArticleToGraphQLArticle(article), nil
}

// AllArticles is the resolver for the allArticles field.
func (r *queryResolver) AllArticles(ctx context.Context) ([]*gqlmodels.Article, error) {
	var contextExecutor = boil.GetContextDB()

	// articles, err := models.Articles(Load(models.Author)).All(ctx, contextExecutor)

	articles, err := models.Articles().All(ctx, contextExecutor)
	if err != nil {
		return nil, err
	}
	// author, err := models.FindAuthor(ctx, contextExecutor, articles.AuthorID)
	// var art []*gqlmodels.Article
	// for idx, a := range articles {
	// 	cnvrttogql.ArticleToGraphQLArticle(a)
	// 	art[idx].Author = cnvrttogql.AuthorToGraphQLAuthor(a.AuthorID)
	// }

	return cnvrttogql.ArticlesToGraphQLArticles(articles), nil
}

// Query returns gqlmodels.QueryResolver implementation.
func (r *Resolver) Query() gqlmodels.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
