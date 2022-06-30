package jsonstrict

import (
	"bytes"
	"encoding/json"
)

func Unmarshal(data []byte, v any) error {
	buf := bytes.NewBuffer(data)
	d := json.NewDecoder(buf)
	d.DisallowUnknownFields()

	return d.Decode(&v)
}
