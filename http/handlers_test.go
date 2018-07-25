package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/golang/mock/gomock"
	"github.com/jsungholee/flock"
	fhttp "github.com/jsungholee/flock/http"
	. "github.com/smartystreets/goconvey/convey"
)

var tweets = []twitter.Tweet{
	twitter.Tweet{
		ID:        1021959695630721024,
		CreatedAt: "Wed Jul 25 03:25:53 +0000 2018",
		Text:      "lorem ipsum #foo bar hello world",
	},
	twitter.Tweet{
		ID:        1021959691864137729,
		CreatedAt: "Wed Jul 25 03:25:52 +0000 2018",
		Text:      "blah blah blah #foo blah blah",
	},
}

var tweetRes = twitter.Search{
	Statuses: tweets,
	Metadata: &twitter.SearchMetadata{
		Count:       2,
		CompletedIn: 0.035,
		Query:       "foo",
	},
}

func Test_HandleSearch(t *testing.T) {
	Convey("Given a valid service implementation and", t, func() {
		mock := gomock.NewController(t)
		svc := flock.NewMockService(mock)

		Convey("When handling a successful HandleSearch request", func() {
			req, err := http.NewRequest("GET", "/test?hashtag=foo", nil)
			if err != nil {
				t.Fatal(err)
			}
			ts := httptest.NewRecorder()

			svc.EXPECT().GetTweets(gomock.Any(), "foo").Return(&tweetRes, nil)
			fhttp.HandleSearch(svc).ServeHTTP(ts, req)

			Convey("An OK status should be returned", func() {
				So(ts.Code, ShouldEqual, http.StatusOK)
			})
		})
	})
}
