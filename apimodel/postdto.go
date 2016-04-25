package apimodel

import (
	"utilities"
	"model"
	"github.com/fatih/structs"
)

type PostDto struct {
	Uuid string
	Title string
	Description string
	Body string
	OwnerId string
}

func FillPostDtoFromMap(m map[string]interface{}) (PostDto, error) {
	i, err := utilities.FillStructFromMap(&PostDto{}, m, false)
	return *i.(*PostDto), err
}

// DTO <--> Model maps

func PostModelFromPostDto(postDto PostDto) model.Post{
	dtoMap := structs.Map(postDto)
	i, _ := utilities.FillStructFromMap(&model.Post{}, dtoMap, true)
	return *i.(*model.Post)
}

func PostDtoFromPostModel(post model.Post) PostDto{
	modelMap := structs.Map(post)
	i, _ := utilities.FillStructFromMap(&PostDto{}, modelMap, true)
	return *i.(*PostDto)
}
