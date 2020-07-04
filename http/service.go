package http

import (
	"net/http"
	gohttp "net/http"
	"time"
)

// API defines the interface to the HTTP service.
type API interface {
	Do(req *gohttp.Request) (*gohttp.Response, error)
}

func New() API {
	return &service{}
}

type service struct {
}

func (s *service) Do(req *gohttp.Request) (*gohttp.Response, error) {
	client := &http.Client{Timeout: time.Second * 10}
	return client.Do(req)
}
