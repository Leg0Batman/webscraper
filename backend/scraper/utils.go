package scraper

import (
	"regexp"
)

func IsValidAPIKey(key string) bool {
	// Add your regex for validating OpenAI API keys
	re := regexp.MustCompile(`your-regex-here`)
	return re.MatchString(key)
}
