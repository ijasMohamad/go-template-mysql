package resolver

import (
	"go-template/gqlmodels"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
     sync.Mutex
     Observers map[string]chan *gqlmodels.Author
     Observers2 map[string]chan *gqlmodels.Article
}
