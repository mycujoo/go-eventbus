package eventbus

import (
	"bytes"
	"encoding/gob"
)

// Encode will encode anything to binary data
func Encode(data interface{}) ([]byte, error) {
	var payload bytes.Buffer
	encoder := gob.NewEncoder(&payload)

	err := encoder.Encode(data)
	if err != nil {
		return []byte{}, err
	}

	return payload.Bytes(), nil
}
