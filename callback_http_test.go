package heartbeat

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCallbackHTTP(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusNoContent)
		cancel()
	}))
	defer srv.Close()

	cb := CallbackHTTP(http.MethodGet, srv.URL)
	err := cb(context.TODO())

	require.NoError(t, err)

	<-ctx.Done()

	require.True(t, errors.Is(ctx.Err(), context.Canceled))
}
