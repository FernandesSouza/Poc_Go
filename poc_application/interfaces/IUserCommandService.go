package application_interfaces

import (
	dtos "poc_api/poc_application/dataTransferObjects/userDtos/request"
)


type IUserCommandService interface{
	Register(user *dtos.UserRegisterRequest) string
}