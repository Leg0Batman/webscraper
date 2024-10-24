package scraper

import (
	"regexp"
)

func IsValidAPIKey(key string) bool {
	// Add your regex for validating OpenAI API keys
	re := regexp.MustCompile(`sk-[a-zA-Z0-9]{48}`)
	return re.MatchString(key)
}
