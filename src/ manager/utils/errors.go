package utils

import (
	"encoding/json"
	"github.com/tomp332/gobrute/src/ manager/managerTypes"
)

func BadRequestError(message string, err error) []byte {
	var s []byte = nil
	if err == nil {
		s, err = json.Marshal(&managerTypes.ServerError{Error: message})
	} else {
		s, err = json.Marshal(&managerTypes.ServerError{Error: err.Error()})
	}
	if err != nil {
		panic(err)
	}
	return s
}
