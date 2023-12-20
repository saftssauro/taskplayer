package helpers

import (
	"encoding/json"
	"io"
)

func DecodeJSONBody(body io.ReadCloser, dto any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&dto)

	return err
}
