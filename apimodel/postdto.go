package apimodel

import (
	"utilities"
	"model"
	"github.com/fatih/structs"
)

type PostDto struct {
	Uuid string `ark-readonly:"true"`
	Title string
	Description string
	Body string
	OwnerId string `ark-readonly:"true"`
}

func (pd PostDto) FillDtoFromMap(m map[string]interface{}) (PostDto, error) {
	i, err := utilities.FillStructFromMap(&PostDto{}, m, false)
	return *i.(*PostDto), err
}

// DTO <--> Model maps

func (pd PostDto) ModelFromDto(postDto PostDto) model.Post{
	dtoMap := structs.Map(postDto)
	i, _ := utilities.FillStructFromMap(&model.Post{}, dtoMap, true)
	return *i.(*model.Post)
}

func (pd PostDto) DtoFromModel(post model.Post) PostDto{
	modelMap := structs.Map(post)
	i, _ := utilities.FillStructFromMap(&PostDto{}, modelMap, true)
	return *i.(*PostDto)
}

func (pd PostDto) DtosFromModels(posts []model.Post) []PostDto {
	var dtoResults []PostDto = make([]PostDto, len(posts))
	for i, v := range posts {
		dto := PostDto{}.DtoFromModel(v)
		dtoResults[i] = dto
	}
	return dtoResults
}
