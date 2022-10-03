package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
	"go-template/pkg/utl"
)

// AuthorNotification is the resolver for the authorNotification field.
func (r *subscriptionResolver) AuthorNotification(ctx context.Context) (<-chan *gqlmodels.Author, error) {
	id := utl.RandomSequence(5)
	event := make(chan *gqlmodels.Author, 1)

	go func() {
		<-ctx.Done()
		r.Lock()
		delete(r.Observers, id)
		r.Unlock()
	}()
	r.Lock()
	r.Observers[id] = event
	r.Unlock()
	fmt.Println("Subscribed to author creation updates!")
	return event, nil
}

// ArticleNotification is the resolver for the articleNotification field.
func (r *subscriptionResolver) ArticleNotification(ctx context.Context) (<-chan *gqlmodels.Article, error) {
	id := utl.RandomSequence(5)
	event := make(chan *gqlmodels.Article, 1)

	go func() {
		<-ctx.Done()
		r.Lock()
		delete(r.Observers2, id)
		r.Unlock()
	}()
	r.Lock()
	r.Observers2[id] = event
	r.Unlock()
	fmt.Println("Subscribed to article creation updates!")
	return event, nil
}

// Subscription returns gqlmodels.SubscriptionResolver implementation.
func (r *Resolver) Subscription() gqlmodels.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
