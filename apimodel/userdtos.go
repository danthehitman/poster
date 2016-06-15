package apimodel

import (
	"utilities"
	"model"
	"github.com/fatih/structs"
)

type RegisterUserDto struct {
	Email string
	FirstName string
	LastName string
	Password string
}

func FillRegisterUserDto(m map[string]interface{}) (RegisterUserDto, error) {
	i, err := utilities.FillStructFromMap(&RegisterUserDto{}, m, false)
	return *i.(*RegisterUserDto), err
}

type UserDto struct {
	Uuid string `ark-readonly:"true"`
	Email string
	FirstName string
	LastName string
}

func FillUserDto(m map[string]interface{}) (UserDto, error) {
	i, err := utilities.FillStructFromMap(&UserDto{}, m, true)
	return *i.(*UserDto), err
}

// DTO <--> Model maps


func UserModelFromRegisterDto(regDto RegisterUserDto) (*model.User, error) {
	dtoMap := structs.Map(regDto)
	i, err := utilities.FillStructFromMap(&model.User{}, dtoMap, true)
	if err != nil{
		return nil, err
	}
	return i.(*model.User), nil
}

func UserDtoFromUserModel(user model.User) (*UserDto, error) {
	modelMap := structs.Map(user)
	i, err := utilities.FillStructFromMap(&UserDto{}, modelMap, true)
	if err != nil{
		return nil, err
	}
	return i.(*UserDto), nil
}