package rq

import (
	"bytes"
	"net/http"
)

func Get(url string) (*Response, error) {
	rs, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return &Response{Http: rs}, nil
}

func Post(url string, contentType string, data []byte) (*Response, error) {
	rs, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return &Response{Http: rs}, nil
}

func NewRequest(method string, url string) (*Request, error) {
	rq, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return &Request{Http: rq, client: &http.Client{}}, nil
}
