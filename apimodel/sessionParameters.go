package apimodel

import (
	"utilities"
)

type SessionParameters struct {
	Email string
	Password string
}

func FillSessionParameters(m map[string]interface{}) (SessionParameters, error) {
	i, err := utilities.FillDto(&SessionParameters{}, m)
	return *i.(*SessionParameters), err
}