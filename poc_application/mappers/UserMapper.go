package mappers

import (
	dtosRequest "poc_api/poc_application/dataTransferObjects/userDtos/request"
	dtosResponse "poc_api/poc_application/dataTransferObjects/userDtos/response"
	application_interfaces "poc_api/poc_application/interfaces"
	"poc_api/poc_domain/entities"
)

type UserMapper struct {
}



func NewUserMapper() application_interfaces.IUserMapper {
	return &UserMapper{}
}

func (um *UserMapper) DtoRegisterToDomain(dto *dtosRequest.UserRegisterRequest) entities.User {
	return entities.User{
		Name:       dto.Name,
		Email:      dto.Email,
		Identifier: dto.Identifier,
	}
}

func (um *UserMapper) DomainToSimpleResponse(domain *entities.User) dtosResponse.UserSimpleResponse {
	return dtosResponse.UserSimpleResponse{
		Name:       domain.Name,
		Email:      domain.Email,
		Identifier: domain.Identifier,
	}
}

func (um *UserMapper) DomainToSimpleResponseList(domain []entities.User) []dtosResponse.UserSimpleResponse {
	var responseList []dtosResponse.UserSimpleResponse

	for _, user := range domain {
		responseList = append(responseList, dtosResponse.UserSimpleResponse(user))
	}

	return responseList
}


func SingleDomainToSimpleResponse(domain *entities.User) *dtosResponse.UserSimpleResponse {
	return &dtosResponse.UserSimpleResponse{
		Name:       domain.Name,
		Email:      domain.Email,
		Identifier: domain.Identifier,
	}
}