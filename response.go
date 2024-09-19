package rq

import (
    "io"
    "net/http"
    "encoding/json"
)

type Response struct {
    Http *http.Response
}

func (r *Response) RawBody() ([]byte, error) {
    body, err := io.ReadAll(r.Http.Body)
    if err != nil {
        return nil, err
    }
    return body, nil
}

func (r *Response) Json(object any) error {
    body, err := r.RawBody()
    if err != nil {
        return err
    }
    err = json.Unmarshal(body, &object)
    if err != nil {
        return err
    }
    return nil
}