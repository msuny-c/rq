package rq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	Http   *http.Request
	client *http.Client
}

func (r *Request) SendBytes(data []byte) (*Response, error) {
	r.Http.Body = io.NopCloser(bytes.NewBuffer(data))
    return r.Send()
}

func (r *Request) SendString(data string) (*Response, error) {
	return r.SendBytes([]byte(data))
}

func (r *Request) SendJson(object any) (*Response, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	return r.SendBytes(data)
}

func (r *Request) Send() (*Response, error) {
	rs, err := r.client.Do(r.Http)
	if err != nil {
		return nil, err
	}
	return &Response{Http: rs}, nil
}
