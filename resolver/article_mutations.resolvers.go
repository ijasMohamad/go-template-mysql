package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/cnvrttogql"
	"strconv"

	null "github.com/volatiletech/null/v8"
)

// CreateArticle is the resolver for the createArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, input gqlmodels.ArticleCreateInput) (*gqlmodels.Article, error) {
	authorId, _ := strconv.Atoi(input.AuthorID)
	article := models.Article{
		Title:    null.StringFrom(input.Title),
		AuthorID: null.IntFrom(authorId),
	}
	// daos....
	newArticle, err := daos.CreateArticle(article, ctx)
	if err != nil {
		return nil, err
	}
	graphArticle := cnvrttogql.ArticleToGraphQLArticle(&newArticle)

	// r.Lock()
	// for _, observer := range r.Observers2 {
	// 	observer <- graphArticle
	// }
	// r.Unlock()

	return graphArticle, nil
}

// UpdateArticle is the resolver for the updateArticle field.
func (r *mutationResolver) UpdateArticle(ctx context.Context, input gqlmodels.ArticleUpdateInput) (*gqlmodels.Article, error) {
	articleID, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, err
	}

	// doas...
	article, _ := daos.FindArticleById(articleID, ctx)

	var a models.Article
	if article != nil {
		a = *article
	}
	if input.Title != nil {
		a.Title = null.StringFrom(*input.Title)
	}

	// doas...
	_, err = daos.UpdateArticle(a, ctx)
	if err != nil {
		return nil, err
	}
	graphArticle := cnvrttogql.ArticleToGraphQLArticle(&a)

	// r.Lock()
	// for _, observer := range r.Observers2 {
	// 	observer <- graphArticle
	// }
	// r.Unlock()

	return graphArticle, nil
}

// DeleteArticle is the resolver for the deleteArticle field.
func (r *mutationResolver) DeleteArticle(ctx context.Context, input *gqlmodels.ArticleDeleteInput) (*gqlmodels.ArticleDeletePayload, error) { //nolint

	articleID, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, err
	}
	// doas...
	article, err := daos.FindArticleById(articleID, ctx)
	if err != nil {
		return nil, err
	}
	// doas...
	_, err = daos.DeleteArticle(*article, ctx)
	if err != nil {
		return nil, err
	}
	return &gqlmodels.ArticleDeletePayload{ID: input.ID}, nil
}

// Mutation returns gqlmodels.MutationResolver implementation.
func (r *Resolver) Mutation() gqlmodels.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
