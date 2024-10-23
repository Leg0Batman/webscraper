package main

import (
	"fmt"
	"os"
	"webscraper/scraper"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <GitHub URL>")
		return
	}
	url := args[1]
	scraper.StartScraping(url)
}
