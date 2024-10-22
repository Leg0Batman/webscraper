package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/gocolly/colly"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run scraper.go <GitHub URL>")
		return
	}
	url := args[1]
	collector := colly.NewCollector()

	// whenever the collector is about to make a new request
	collector.OnRequest(func(r *colly.Request) {
		// print the url of that request
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
	})
	collector.OnHTML("body", func(e *colly.HTMLElement) {
		bodyText := e.Text
		if containsOpenAIKey(bodyText) {
			fmt.Println("Alarm! Found an OpenAI API key on", e.Request.URL)
		}
	})
	collector.Visit(url)
}


func containsOpenAIKey(text string) bool {
	re := regexp.MustCompile(`sk-[a-zA-Z0-9]{48}`) // OpenAI API keys are 48 characters long with the prefix "sk-"
	return re.MatchString(text)
}
