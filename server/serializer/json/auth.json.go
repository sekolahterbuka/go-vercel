package json

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/sekolahkita/go-api/server/model"
)

type Auth struct{}

func (r *Auth) Decode(input []byte) (*model.Auth, error) {
	authJson := &model.Auth{}
	if err := json.Unmarshal(input, authJson); err != nil {
		return nil, errors.Wrap(err, "serializer.Auth.Decode")
	}
	return authJson, nil
}

func (r *Auth) Encode(input *model.Auth) ([]byte, error) {
	authRaw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Auth.Encode")
	}
	return authRaw, nil
}
