package heartbeat

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
)

func CallbackHTTP(method, url string) Callback {
	c := cleanhttp.DefaultClient()

	return func(ctx context.Context) error {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return err
		}
		res, err := c.Do(req.WithContext(ctx))
		if err != nil {
			return err
		}
		defer res.Body.Close()

		_, err = io.Copy(ioutil.Discard, res.Body)
		if err != nil {
			return err
		}

		return nil
	}
}
