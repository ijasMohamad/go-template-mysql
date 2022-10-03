package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/pkg/utl/cnvrttogql"
)

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, id int) (*gqlmodels.Article, error) {
	// doas...
	article, err := daos.FindArticleById(id, ctx)
	if err != nil {
		return nil, err
	}
	return cnvrttogql.ArticleToGraphQLArticle(article), nil
}

// AllArticles is the resolver for the allArticles field.
func (r *queryResolver) AllArticles(ctx context.Context) (*gqlmodels.ArticlesPayload, error) {
	// daos...
	articles, count, err := daos.FindAllArticlesWithCount(ctx)
	if err != nil {
		return nil, err
	}
	graphqlArticles := cnvrttogql.ArticlesToGraphQLArticles(articles)
	return &gqlmodels.ArticlesPayload{
		Articles: graphqlArticles,
		Total:    count}, nil
}

// Query returns gqlmodels.QueryResolver implementation.
func (r *Resolver) Query() gqlmodels.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
