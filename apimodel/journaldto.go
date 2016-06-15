package apimodel

import (
	"utilities"
	"model"
	"github.com/fatih/structs"
)

type CreateJournalDto struct {
	Title string
	Description string
	IsPublic bool
	OwnerId string `ark-readonly:"true"`
}

type JournalDto struct {
	Uuid string `ark-readonly:"true"`
	Title string
	Description string
	IsPublic bool
	OwnerId string `ark-readonly:"true"`
}

func (jd JournalDto) FillDtoFromMap(m map[string]interface{}) (JournalDto, error) {
	i, err := utilities.FillStructFromMap(&JournalDto{}, m, false)
	return *i.(*JournalDto), err
}

// DTO <--> Model maps

func (jd JournalDto) ModelFromDto(dto JournalDto) model.Journal{
	dtoMap := structs.Map(dto)
	i, _ := utilities.FillStructFromMap(&model.Journal{}, dtoMap, true)
	return *i.(*model.Journal)
}

func (jd JournalDto) DtoFromModel(entity model.Journal) JournalDto{
	modelMap := structs.Map(entity)
	i, _ := utilities.FillStructFromMap(&JournalDto{}, modelMap, true)
	return *i.(*JournalDto)
}

func (jd JournalDto) DtosFromModels(entities []model.Journal) []JournalDto {
	var dtoResults []JournalDto = make([]JournalDto, len(entities))
	for i, v := range entities {
		dto := jd.DtoFromModel(v)
		dtoResults[i] = dto
	}
	return dtoResults
}
