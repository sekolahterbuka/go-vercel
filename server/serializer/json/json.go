package json

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func Decode[K comparable](input []byte) (*K, error) {
	var data K
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, errors.Wrap(err, "serializer.Decode")
	}
	return &data, nil
}

func Encode[K comparable](input *K) ([]byte, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Encode")
	}
	return data, nil
}
