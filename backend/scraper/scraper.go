package scraper

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func StartScraping(url string) (string, error) {
	collector := colly.NewCollector()

	// Rate limiting and caching
	rateLimiter := NewRateLimiter()
	cache := NewCache()

	var result string
	var scrapeError error

	collector.OnRequest(func(r *colly.Request) {
		rateLimiter.Wait()
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
		cache.Set(r.Request.URL.String(), r.Body)
		// Example: Adjust the selector to find the API key
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
		if err != nil {
			scrapeError = err
			return
		}
		apiKey := doc.Find("meta[name='api-key']").AttrOr("content", "")
		if apiKey == "" {
			scrapeError = fmt.Errorf("API key not found")
		} else {
			result = apiKey
		}
	})

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error:", e)
		scrapeError = e
	})

	err := collector.Visit(url)
	if err != nil {
		return "", err
	}

	if scrapeError != nil {
		return "", scrapeError
	}

	return result, nil
}
