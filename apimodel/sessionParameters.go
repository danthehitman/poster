package apimodel

import "utilities"

type SessionParameters struct {
	Email string
	Password string
}

func FillSessionParametersStruct(m map[string]interface{}) (*SessionParameters, error) {
	s := &SessionParameters{}
	for k, v := range m {
		err := utilities.SetField(s, k, v)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}