package flock

import (
	"context"

	"github.com/dghubble/go-twitter/twitter"
)

// Service defines the public interface that an api needs to expose
type Service interface {
	GetTweets(ctx context.Context, searchTerm string) (*twitter.Search, error)
}
