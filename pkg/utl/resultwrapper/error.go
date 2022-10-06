package resultwrapper

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func HandleGraphQLError(errMsg string) graphql.ResponseHandler {
     return func(ctx context.Context) *graphql.Response {
          return &graphql.Response{
               Errors: gqlerror.List{gqlerror.Errorf(errMsg)},
          }
     }
}