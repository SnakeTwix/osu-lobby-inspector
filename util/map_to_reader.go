package util

import (
	"bytes"
	"encoding/json"
	"io"
)

func MapToReader(content map[string]any) (io.Reader, error) {
	jsonContent, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonContent), nil
}
