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

func FillPostDto(m map[string]interface{}) (PostDto, error) {
	//s := &PostDto{}
	//i, err := utilities.FillDto(s, m)
	//s = i.(*PostDto)
	//return *s, err
	i, err := utilities.FillStructFromMap(&PostDto{}, m, false)
	return *i.(*PostDto), err
}

// DTO <--> Model maps

func PostFromPostDto(postDto PostDto) model.Post{
	dtoMap := structs.Map(postDto)
	i, _ := utilities.FillStructFromMap(&model.Post{}, dtoMap, true)
	return *i.(*model.Post)
	//post := model.Post{
	//	Description: postDto.Description,
	//	OwnerId: postDto.OwnerId,
	//	Title: postDto.Title,
	//	Body: postDto.Body,
	//	Uuid: postDto.Uuid,
	//}
	//return post
}

func PostDtoFromPost(post model.Post) PostDto{
	modelMap := structs.Map(post)
	i, _ := utilities.FillStructFromMap(&PostDto{}, modelMap, true)
	return *i.(*PostDto)
	//postDto := PostDto{
	//	Description: post.Description,
	//	Title: post.Title,
	//	Body: post.Body,
	//	OwnerId: post.OwnerId,
	//	Uuid: post.Uuid,
	//}
	//return postDto
}
