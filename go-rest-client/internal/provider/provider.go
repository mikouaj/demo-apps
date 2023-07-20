// Package provider provides data providers
package provider

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/mikouaj/go-rest-client/internal/api"
)

type DataProvider interface {
	GetData(ctx context.Context) (interface{}, error)
}

type httpDataProvider struct {
	appName string
	url     string
}

type localDataProvider struct {
	appName string
}

func NewLocalDataProvider(appName string) DataProvider {
	return &localDataProvider{
		appName: appName,
	}
}

func (p *localDataProvider) GetData(context.Context) (interface{}, error) {
	books := api.Books{
		{
			Title:    "Ubik",
			Author:   "Philip K. Dick",
			Category: "Science Fiction",
		},
		{
			Title:    "Enders game",
			Author:   "Orson Scott Card",
			Category: "Science Fiction",
		},
	}
	return &api.Data{
		Source: p.appName,
		Data:   books,
	}, nil
}

func NewHTTPDataProvider(appName, url string) DataProvider {
	return &httpDataProvider{
		appName: appName,
		url:     url,
	}
}

func (p *httpDataProvider) GetData(ctx context.Context) (interface{}, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var dataJSON map[string]interface{}
	if err := json.Unmarshal(data, &dataJSON); err != nil {
		return nil, err
	}
	return &api.Data{
		Source: p.url,
		Data:   dataJSON,
	}, nil
}
