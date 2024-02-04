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
	"sync" // Add this import for synchronization
	"time"
)

// var newsapi *news.Client
var tpl = template.Must(template.ParseFiles("index.html"))

// Cache to store previously fetched results
var resultCache = make(map[string]*news.Results)
var cacheMutex sync.Mutex // Mutex for synchronization when accessing the cache
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

		// Check if the result is already in the cache
		cacheKey := searchQuery + "_" + page
		cacheMutex.Lock()
		cachedResult, exists := resultCache[cacheKey]
		cacheMutex.Unlock()

		if exists {
			// If result is in the cache, use it directly
			search := &Search{
				Query:      searchQuery,
				NextPage:   nextPage,
				TotalPages: int(math.Ceil(float64(cachedResult.TotalResults) / float64(newsapi.PageSize))),
				Results:    cachedResult,
			}

			renderSearchTemplate(w, search)
			return
		}

		// If result is not in the cache, fetch it from the API
		results, err := newsapi.FetchEverything(searchQuery, page)
		if err != nil {
			log.Printf("Error fetching news: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Cache the result for future use
		cacheMutex.Lock()
		resultCache[cacheKey] = results
		cacheMutex.Unlock()

		// Render the search template
		nextPage, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		search := &Search{
			Query:      searchQuery,
			NextPage:   nextPage,
			TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(newsapi.PageSize))),
			Results:    results,
		}

		if ok := !search.IsLastPage(); ok {
			search.NextPage++
		}

		buf := &bytes.Buffer{}
		err = tpl.Execute(buf, search)
		if err != nil {
			log.Printf("Error rendering search template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	}
}

// Helper function to render the search template
func renderSearchTemplate(w http.ResponseWriter, search *Search) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, search)
	if err != nil {
		log.Printf("Error rendering search template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
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
