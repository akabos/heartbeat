package heartbeat

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHeartbeat_Send(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*9)
	defer cancel()

	c := 0
	d := 0

	h := New(time.Millisecond*1, func(ctx context.Context) error {
		c++
		return nil
	})

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				d++
				_ = h.Send(ctx)
			}
		}
	}( ctx)

	<-ctx.Done()

	require.Greater(t, d, c)
	require.Equal(t, 10, c)
}
