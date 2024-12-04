package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct {
}

// 最簡單的封裝
func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
