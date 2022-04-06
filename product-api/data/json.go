package data

import (
	"encoding/json"
	"io"
)

func FromJson(s interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(s)
}

func ToJson(s interface{}, rw io.Writer) error {
	encoder := json.NewEncoder(rw)
	return encoder.Encode(s)
}
