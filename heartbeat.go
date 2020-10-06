package heartbeat

import (
	"context"
	"errors"
	"time"

	"golang.org/x/time/rate"
)

var ErrRateLimitExceeded = errors.New("rate limit exceeded")

type Callback func(ctx context.Context) error

func New(every time.Duration, cb Callback) *Heartbeat {
	return &Heartbeat{
		every: every,
		cb:    cb,
		lim:   rate.NewLimiter(rate.Every(every), 1),
	}
}

type Heartbeat struct {
	every time.Duration
	cb    Callback
	lim   *rate.Limiter
}

func (h *Heartbeat) Send(ctx context.Context) error {
	if !h.lim.Allow() {
		return ErrRateLimitExceeded
	}
	ctx, cancel := context.WithTimeout(ctx, h.every)
	defer cancel()
	return h.cb(ctx)
}
