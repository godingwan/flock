package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jsungholee/flock/http"
	"github.com/jsungholee/flock/service"
)

func main() {
	host := getEnvString("HTTP_HOST", "0.0.0.0")
	port, err := strconv.Atoi(getEnvString("HTTP_PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to parse port integer value", err)
	}

	APIKey := getEnvString("TWITTER_API_KEY", "44IAzOwv8EkK6VCSskQV8Aenq")
	APISec := getEnvString("TWITTER_API_SECRET", "x1fUSFYg4pMGwM7Bl20II4GFS4UH4KS9k1mYETF4BHDc3vjeum")
	at := getEnvString("TWITTER_ACCESS_TOKEN", "909040847227965440-0IdWVhlWwCEVX89CZIqfA5T3o593Xq4")
	ats := getEnvString("TWITTER_ACCESS_TOKEN_SECRET", "VAIv9zjbhZzgW5IvhfI6EQZr6p0tkQUDluDSzpxsq8sMa")

	config := oauth1.NewConfig(APIKey, APISec)
	token := oauth1.NewToken(at, ats)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	svc := service.DefaultService{
		TClient: client,
	}

	log.Printf("Server starting on %s:%d", host, port)
	server := http.NewServer(host, uint(port), http.BuildRouter(svc))
	log.Printf("Server Failed: %s", server.Start())
}

func getEnvString(env, def string) string {
	val := os.Getenv(env)
	if val == "" {
		return def
	}
	return val
}
