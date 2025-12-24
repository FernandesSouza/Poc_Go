package application_interfaces

import (
	requestDtos "poc_api/poc_application/dataTransferObjects/userDtos/request"
	responseDtos "poc_api/poc_application/dataTransferObjects/userDtos/response"
	"poc_api/poc_domain/entities"
)

type IUserMapper interface {
	DtoRegisterToDomain(dtoRegister *requestDtos.UserRegisterRequest) entities.User
	DomainToSimpleResponse(domain * entities.User) responseDtos.UserSimpleResponse
	DomainToSimpleResponseList(domain []entities.User) []responseDtos.UserSimpleResponse
}
