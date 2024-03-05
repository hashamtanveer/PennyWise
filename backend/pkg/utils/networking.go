package utils

import (
	"github.com/bytedance/sonic"
	"io"
    "net/http"
    "errors"
)

func BindJsonBody(r *http.Request, to interface{}) error {
    if r.Header.Get("Content-Type") != "application/json" {
        return errors.New("Content-Type not json")
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        return err
    }

    return sonic.Unmarshal(body, to)
}
