package util

import (
	"context"
	"time"
)

// Ticker repeatedly calls cb with a delay in between calls. It stops doing This
// When a element is sent to the done channel.
func Ticker(cb func(), delay time.Duration, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			cb()
			time.Sleep(delay)
		}
	}
}
