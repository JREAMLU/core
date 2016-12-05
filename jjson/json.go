package jjson

import (
	"bytes"
	"encoding/json"
)

// PrettyJSON pretty json
func PrettyJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

// Minify minify
func Minify(input string) (string, error) {
	var out bytes.Buffer
	reader := bytes.NewBufferString(input)
	err := WriteMinifiedTo(&out, reader)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
