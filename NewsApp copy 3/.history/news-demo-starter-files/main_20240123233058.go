package main

import (
	"math"

	"github.com/freshman-tech/news-demo-starter-files/news"

	"bytes"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"context"
	"github.com/go-redis/redis/v8"
)

// var newsapi *news.Client
var tpl = template.Must(template.ParseFiles("index.html"))

// Redis client to connect to the Redis server
var redisClient *redis.Client

type Search struct {
	Query        string
	NextPage     int
	TotalPages   int
	Results      *news.Results
	Loading      bool // Add this field for indicating whether data is being loaded
	Error        bool // Add this field for indicating if an error occurred
	ErrorMessage string
}

func (s *Search) IsLastPage() bool {
	return s.NextPage >= s.TotalPages
}

func (s *Search) CurrentPage() int {
	if s.NextPage == 1 {
		return s.NextPage
	}

	return s.NextPage - 1
}

func (s *Search) PreviousPage() int {
	return s.CurrentPage() - 1
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		log.Printf("Error rendering index template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func searchHandler(newsapi *news.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("Error parsing URL: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		params := u.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}

		// Check if the result is already in Redis cache
		cacheKey := searchQuery + "_" + page
		cachedResult, err := getFromCache(cacheKey)
		if err == nil && cachedResult != nil {
			// If result is in the cache, use it directly
			renderSearchTemplate(w, searchQuery, page, cachedResult)
			return
		}

		// If result is not in the cache, fetch it from the API
		results, err := newsapi.FetchEverything(searchQuery, page)
		if err != nil {
			log.Printf("Error fetching news: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Cache the result in Redis for future use
		err = setToCache(cacheKey, results)
		if err != nil {
			log.Printf("Error caching result in Redis: %v", err)
		}

		// Render the search template
		renderSearchTemplate(w, searchQuery, page, results)
	}
}

// Helper function to get data from Redis cache
func getFromCache(key string) (*news.Results, error) {
	ctx := context.Background()
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var result news.Results
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func init() {
	// Initialize Redis client in the init function
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Change this to your Redis server address
		Password: "",               // No password for local Redis server
		DB:       0,                // Default DB
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}
	newsapi := news.NewClient(myClient, apiKey, 20)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/search", searchHandler(newsapi))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
