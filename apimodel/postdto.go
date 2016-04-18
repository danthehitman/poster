package apimodel

import "utilities"

type PostDto struct {
	Uuid string
	Title string
	Description string
	OwnerId string
}

func FillPostDto(m map[string]interface{}) (*PostDto, error) {
	s := &PostDto{}
	for k, v := range m {
		err := utilities.SetField(s, k, v)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}
