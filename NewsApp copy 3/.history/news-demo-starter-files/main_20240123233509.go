package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/freshman-tech/news-demo-starter-files/news"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

// App struct to hold components
type App struct {
	NewsAPI     *news.Client
	Template    *template.Template
	RedisClient *redis.Client
}

// Search struct
type Search struct {
	Query        string
	NextPage     int
	TotalPages   int
	Results      *news.Results
	Loading      bool
	Error        bool
	ErrorMessage string
}

// NewApp initializes the App struct
func NewApp() (*App, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Change this to your Redis server address
		Password: "",               // No password for local Redis server
		DB:       0,                // Default DB
	})

	// Initialize News API client
	myClient := &http.Client{Timeout: 10 * time.Second}
	newsAPIKey := os.Getenv("NEWS_API_KEY")
	if newsAPIKey == "" {
		return nil, fmt.Errorf("Env: NEWS_API_KEY must be set")
	}
	newsAPI := news.NewClient(myClient, newsAPIKey, 20)

	// Load template
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		return nil, err
	}

	return &App{
		NewsAPI:     newsAPI,
		Template:    tpl,
		RedisClient: redisClient,
	}, nil
}

// searchHandler handles search requests
func (app *App) searchHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	searchQuery := params.Get("q")
	page := params.Get("page")
	if page == "" {
		page = "1"
	}

	// Check if the result is already in Redis cache
	cacheKey := searchQuery + "_" + page
	cachedResult, err := app.getFromCache(cacheKey)
	if err == nil && cachedResult != nil {
		// If result is in the cache, use it directly
		app.renderSearchTemplate(w, searchQuery, page, cachedResult)
		return
	}

	// If result is not in the cache, fetch it from the API
	results, err := app.NewsAPI.FetchEverything(searchQuery, page)
	if err != nil {
		log.Printf("Error fetching news: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Cache the result in Redis for future use
	err = app.setToCache(cacheKey, results)
	if err != nil {
		log.Printf("Error caching result in Redis: %v", err)
	}

	// Render the search template
	app.renderSearchTemplate(w, searchQuery, page, results)
}

// Helper function to get data from Redis cache
func (app *App) getFromCache(key string) (*news.Results, error) {
	ctx := context.Background()
	val, err := app.RedisClient.Get(ctx, key).Result()
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

// Helper function to set data to Redis cache
func (app *App) setToCache(key string, value interface{}) error {
	ctx := context.Background()
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return app.RedisClient.Set(ctx, key, jsonValue, 0).Err()
}

// Helper function to render the search template
func (app *App) renderSearchTemplate(w http.ResponseWriter, searchQuery, page string, results *news.Results) {
	nextPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	search := &Search{
		Query:      searchQuery,
		NextPage:   nextPage,
		TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(app.NewsAPI.PageSize))),
		Results:    results,
	}

	if ok := !search.IsLastPage(); ok {
		search.NextPage++
	}

	buf := &bytes.Buffer{}
	err = app.Template.Execute(buf, search)
	if err != nil {
		log.Printf("Error rendering search template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func (app *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := app.Template.Execute(buf, nil)
	if err != nil {
		log.Printf("Error rendering index template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/search", app.searchHandler)
	mux.HandleFunc("/", app.indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
