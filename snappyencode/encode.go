// Created by davidterranova on 27/09/2017.

package snappyencode

import (
	"bytes"

	"encoding/gob"

	snappystream "github.com/mreiferson/go-snappystream"
)

func Encode(o interface{}) ([]byte, error) {
	var buffer bytes.Buffer

	fs := snappystream.NewBufferedWriter(&buffer)
	encoder := gob.NewEncoder(fs)

	err := encoder.Encode(o)
	if err != nil {
		return nil, err
	}
	err = fs.Close()
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Decode(data []byte, o interface{}) error {
	buffer := bytes.Buffer{}
	buffer.Write(data)

	fs := snappystream.NewReader(&buffer, true)
	decoder := gob.NewDecoder(fs)

	return decoder.Decode(o)
}
