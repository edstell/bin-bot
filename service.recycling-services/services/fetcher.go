package services

import (
	"context"
	"io/ioutil"
	"net/http"

	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
)

// Fetcher implementations should retrieve the latest services available for
// the propertyID provided.
type Fetcher interface {
	Fetch(context.Context, string) ([]recyclingservices.Service, error)
}

type FetcherFunc func(context.Context, string) ([]recyclingservices.Service, error)

func (f FetcherFunc) Fetch(ctx context.Context, propertyID string) ([]recyclingservices.Service, error) {
	return f(ctx, propertyID)
}

// WebScraper returns a Fetcher implementation which will fetch the latest
// service information by scraping the recycling services website.
func WebScraper(client *http.Client, parser Parser, baseURL string) Fetcher {
	return FetcherFunc(func(ctx context.Context, propertyID string) ([]recyclingservices.Service, error) {
		req, err := http.NewRequest("GET", baseURL+"/"+propertyID, nil)
		if err != nil {
			return nil, err
		}
		req = req.WithContext(ctx)

		rsp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer rsp.Body.Close()

		bytes, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return nil, err
		}

		return parser.Parse(bytes)
	})
}