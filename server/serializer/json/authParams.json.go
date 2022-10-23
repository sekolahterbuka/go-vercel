package json

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/sekolahkita/go-api/server/model"
)

type RegisterParams struct{}

func (r *RegisterParams) Decode(input []byte) (*model.RegisterParams, error) {
	registerJson := &model.RegisterParams{}
	if err := json.Unmarshal(input, registerJson); err != nil {
		return nil, errors.Wrap(err, "serializer.RegisterParams.Decode")
	}
	return registerJson, nil
}

func (r *RegisterParams) Encode(input *model.RegisterParams) ([]byte, error) {
	registerRaw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.RegisterParams.Encode")
	}
	return registerRaw, nil
}
