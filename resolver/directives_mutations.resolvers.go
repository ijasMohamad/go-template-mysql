package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-template/daos"
	"go-template/gqlmodels"
	"strconv"
)

// DeleteAuthor is the resolver for the deleteAuthor field.
func (r *mutationResolver) DeleteAuthor(ctx context.Context, input *gqlmodels.AuthorDeleteInput) (*gqlmodels.AuthorDeletePayload, error) {
	//nolint
	authorID, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, err
	}
	// daos..
	author, err := daos.FindAuthorById(authorID, ctx)
	if err != nil {
		return nil, err
	}
	_, err = daos.DeleteAuthor(*author, ctx)
	if err != nil {
		return nil, err
	}
	return &gqlmodels.AuthorDeletePayload{ID: input.ID}, nil
}
