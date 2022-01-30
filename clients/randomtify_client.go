package clients

import (
	"fmt"
	"net/http"
	"net/url"
	"randomtify_app/services"
	"time"
)

type RandomtifyClient interface {
	GetAlphabets() (resp *http.Response, err error)
	GetArtist(name string) (resp *http.Response, err error)
	GetRandomArtist(query, alphabet, charsAmount string) (resp *http.Response, err error)
}

type randomtifyClient struct {
	client *http.Client
	config *services.Config
}

func GetRandomtifyClient() RandomtifyClient {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	return &randomtifyClient{
		client: client,
		config: services.GetConfig(),
	}
}

func (rc *randomtifyClient) GetAlphabets() (resp *http.Response, err error) {
	return rc.client.Get(rc.config.RandomtifyAppUrl + rc.config.SearchPath + rc.config.AlphabetsPath)
}

func (rc *randomtifyClient) GetArtist(name string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", rc.config.RandomtifyAppUrl+rc.config.ArtistsPath, nil)
	if name != "" {
		q := &req.URL.Path
		*q = *q + fmt.Sprintf("/%s", name)
	}
	return rc.client.Do(req)
}

func (rc *randomtifyClient) GetRandomArtist(query, alphabet, charsAmount string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", rc.config.RandomtifyAppUrl+rc.config.SearchPath, nil)
	q := req.URL.Query()
	addUrlQueryParam(&q, "query", query)
	addUrlQueryParam(&q, "alphabet", alphabet)
	addUrlQueryParam(&q, "charsAmount", charsAmount)
	req.URL.RawQuery = q.Encode()
	return rc.client.Do(req)
}

func addUrlQueryParam(query *url.Values, paramName, value string) {
	if value != "" {
		(*query).Add(paramName, value)
	}
}
