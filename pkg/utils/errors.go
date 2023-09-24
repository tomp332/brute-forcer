package utils

import (
	"encoding/json"
	"github.com/tomp332/gobrute/pkg/types"
)

func BadRequestError(message string, err error) []byte {
	var s []byte = nil
	if err == nil {
		s, err = json.Marshal(&types.ServerError{Error: message})
	} else {
		s, err = json.Marshal(&types.ServerError{Error: err.Error()})
	}
	if err != nil {
		panic(err)
	}
	return s
}
