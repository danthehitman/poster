package apimodel

import (
	"utilities"
)

type PostDto struct {
	Uuid string
	Title string
	Description string
	Body string
	OwnerId string
}

func FillPostDto(m map[string]interface{}) (PostDto, error) {
	//s := &PostDto{}
	//i, err := utilities.FillDto(s, m)
	//s = i.(*PostDto)
	//return *s, err
	i, err := utilities.FillDto(&PostDto{}, m)
	return *i.(*PostDto), err
}
