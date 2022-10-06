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

	null "github.com/volatiletech/null/v8"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*gqlmodels.LoginResponse, error) {
	author, err := daos.FindAuthorByUsername(username, ctx)
	if err != nil {
		return nil, err
	}
	// loading configurations
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	// creating new secure and token generation service
	sec := service.Secure(cfg)
	tg, err := service.JWT(cfg)
	if err != nil {
		return nil, fmt.Errorf("error in creating auth services")
	}
	if !author.Password.Valid || (!sec.HashMatchesPassword(author.Password.String, password)) {
		return nil, fmt.Errorf("username or password does not exist")
	}
	if !author.Active.Bool {
		return nil, fmt.Errorf("unathorized author")
	}

	token, err := tg.GenerateToken(author)
	if err != nil {
		return nil, fmt.Errorf("error in generating token: %s", err)
	}
	refreshToken := sec.Token(token)
	author.Token = null.StringFrom(refreshToken)
	_, err = daos.UpdateAuthor(*author, ctx)
	if err != nil {
		return nil, fmt.Errorf("error in updating author")
	}
	return &gqlmodels.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
