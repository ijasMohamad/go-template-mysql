package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/pkg/utl/cnvrttogql"
)

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id int) (*gqlmodels.Author, error) {
	author, err := daos.FindAuthorById(id, ctx)
	if err != nil {
		return nil, err
	}
	return cnvrttogql.AuthorToGraphQLAuthor(author), nil
}

// AllAuthors is the resolver for the allAuthors field.
func (r *queryResolver) AllAuthors(ctx context.Context) (*gqlmodels.AuthorsPayload, error) {
	authors, count, err := daos.FindAllAuthorsWithCount(ctx)
	if err != nil {
		return nil, err
	}
	graphqlAuthor := cnvrttogql.AuthorsToGraphQLAuthors(authors)
	return &gqlmodels.AuthorsPayload{
		Authors: graphqlAuthor,
		Total:   count}, nil
}
