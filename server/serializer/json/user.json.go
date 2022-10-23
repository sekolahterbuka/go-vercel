package json

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/sekolahkita/go-api/server/model"
)

type User struct{}

func (r *User) Decode(input []byte) (*model.User, error) {
	redirect := &model.User{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.User.Decode")
	}
	return redirect, nil
}

func (r *User) Encode(input *model.User) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.User.Encode")
	}
	return rawMsg, nil
}
