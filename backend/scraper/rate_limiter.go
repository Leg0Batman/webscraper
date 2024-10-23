package scraper

import (
	"context"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiter *rate.Limiter
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(1, 5), // 1 request per second with a burst of 5
	}
}

func (r *RateLimiter) Wait() {
	r.limiter.Wait(context.Background())
}
