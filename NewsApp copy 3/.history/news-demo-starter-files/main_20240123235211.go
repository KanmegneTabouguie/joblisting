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
)

// var newsapi *news.Client
var tpl = template.Must(template.ParseFiles("index.html"))

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

		results, err := newsapi.FetchEverything(searchQuery, page)
		if err != nil {
			log.Printf("Error fetching news: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			rfunc searchHandler(newsapi *news.Client) http.HandlerFunc {
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
			
					cachedResults, cached := searchCache.Get(searchQuery)
					if cached {
						log.Printf("Using cached results for query: %s", searchQuery)
						results := cachedResults.(*news.Results)
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
						return
					}
			
					results, err := newsapi.FetchEverything(searchQuery, page)
					if err != nil {
						log.Printf("Error fetching news: %v", err)
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
						return
					}
			
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
			
					searchCache.Set(searchQuery, results, time.Minute) // Set the cache to expire in 1 minute
			
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
			
				// Create a cache
				searchCache := &sync.Map{}
			
				mux := http.NewServeMux()
			
				fs := http.FileServer(http.Dir("assets"))
				mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
			
				// Add searchHandler to the mux
				mux.HandleFunc("/search", searchHandler(searchCache))
				mux.HandleFunc("/", indexHandler)
			
				http.ListenAndServe(":"+port, mux)
			}
			
