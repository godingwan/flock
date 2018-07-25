package service

import (
	"context"
	"log"

	"github.com/pkg/errors"

	"github.com/dghubble/go-twitter/twitter"
)

// DefaultService is the default implementation of the api
type DefaultService struct {
	TClient *twitter.Client
}

// GetTweets will reach out to the twitter api to get tweets
func (s DefaultService) GetTweets(ctx context.Context, searchTerm string) (*twitter.Search, error) {
	// Search Tweets
	tweets, _, err := s.TClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: searchTerm,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get tweets")
	}

	log.Println("Tweets are: ", tweets)
	// fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)
	return tweets, nil
}
