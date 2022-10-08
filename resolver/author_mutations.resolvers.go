package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/internal/config"
	"go-template/internal/service"
	"go-template/models"
	"go-template/pkg/utl/cnvrttogql"
	"strconv"

	null "github.com/volatiletech/null/v8"
)

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, input gqlmodels.AuthorCreateInput) (*gqlmodels.Author, error) {
	active := null.NewBool(false, false)
	if input.Active != nil {
		active = null.BoolFrom(*input.Active)
	}
	author := models.Author{
		FirstName: null.StringFrom(input.FirstName),
		LastName:  null.StringFrom(input.LastName),
		Username:  null.StringFrom(input.Username),
		Password:  null.StringFrom(input.Password),
		Active:    active,
	}
	cfg, err := config.Load()
	if err != nil {

		fmt.Println("Error in loading config")
		return nil, fmt.Errorf("Error in loading config")
	}
	sec := service.Secure(cfg)
	fmt.Println("INPUT Password: ", author.Password.String)
	if author.Password.String != "" {
		author.Password = null.StringFrom(sec.Hash(author.Password.String))
	}
	fmt.Println("Hashed password: ", author.Password)
	// daos..
	newAuthor, err := daos.CreateAuthor(author, ctx)
	if err != nil {
		fmt.Println("Error in creating author in doas", err)
		fmt.Println("New Author: ", newAuthor)
		return nil, err
	}
	graphqlAuthor := cnvrttogql.AuthorToGraphQLAuthor(&newAuthor)

	fmt.Println("GRAPHQL AUTHOR: ", graphqlAuthor)

	graphqlAuthor.Password = nil

	return graphqlAuthor, nil
}

// UpdateAuthor is the resolver for the updateAuthor field.
func (r *mutationResolver) UpdateAuthor(ctx context.Context, input *gqlmodels.AuthorUpdateInput) (*gqlmodels.Author, error) {

	fmt.Println("INPUT ID: ", input.ID)
	authorID, err := strconv.Atoi(input.ID)
	if err != nil {
		fmt.Println("Error in resolver", err)
		return nil, err
	}

	// daos..
	author, err := daos.FindAuthorById(authorID, ctx)
	if err != nil {
		return nil, err
	}
	var a models.Author
	if author != nil {
		a = *author
	}
	if input.FirstName != nil {
		a.FirstName = null.StringFrom(*input.FirstName)
	}
	if input.LastName != nil {
		a.LastName = null.StringFrom(*input.LastName)
	}
	if input.Username != nil {
		a.Username = null.StringFrom(*input.Username)
	}
	if input.Active != nil {
		a.Active = null.BoolFrom(*input.Active)
	}
	if input.Password != nil {
		cfg, err := config.Load()
		if err != nil {
			return nil, fmt.Errorf("Error in loading config")
		}
		sec := service.Secure(cfg)
		a.Password = null.StringFrom(sec.Hash(*input.Password))
	}
	// daos...
	_, err = daos.UpdateAuthor(a, ctx)
	if err != nil {
		fmt.Println("Error while updating author", err)
		return nil, err
	}
	graphqlAuthor := cnvrttogql.AuthorToGraphQLAuthor(&a)
	return graphqlAuthor, nil
}

// DeleteAuthor is the resolver for the deleteAuthor field.
func (r *mutationResolver) DeleteAuthor(ctx context.Context, input *gqlmodels.AuthorDeleteInput) (*gqlmodels.AuthorDeletePayload, error) { //nolint
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
