package scraper

import (
    "fmt"
    "github.com/gocolly/colly"
    "webscraper/scraper/utils"
)

func StartScraping(url string) {
    collector := colly.NewCollector()

    // Rate limiting and caching
    rateLimiter := utils.NewRateLimiter()
    cache := utils.NewCache()

    collector.OnRequest(func(r *colly.Request) {
        rateLimiter.Wait()
        fmt.Println("Visiting", r.URL)
    })

    collector.OnResponse(func(r *colly.Response) {
        fmt.Println("Got a response from", r.Request.URL)
        cache.Set(r.Request.URL.String(), r.Body)
    })

    collector.OnError(func(r *colly.Response, e error) {
        fmt.Println("Error:", e)
    })

    collector.Visit(url)
}