package eventbus

import (
	"bytes"
	"encoding/gob"
)

// Decode will decode a binary data to a target struct
func Decode(message []byte, target interface{}) error {
	decoder := gob.NewDecoder(bytes.NewBuffer(message))
	err := decoder.Decode(target)
	if err != nil {
		return err
	}

	return nil
}
