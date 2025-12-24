package services

import (
	dtos "poc_api/poc_application/dataTransferObjects/userDtos/response"
	application_interfaces "poc_api/poc_application/interfaces"
	repository "poc_api/poc_infra/interfaces"
)

type UserQueryService struct {
	repository repository.IUserRepository
	mapper application_interfaces.IUserMapper
}

func NewUserQueryService(
	userRepository repository.IUserRepository,
	userMapper application_interfaces.IUserMapper,
) application_interfaces.IUserQueryService {
	return &UserQueryService{
		repository: userRepository,
		mapper:  userMapper,
	}
}

func (u *UserQueryService) FindAll() []dtos.UserSimpleResponse {
	 var resultRepository = u.repository.FindAllUsers()

	 if len(resultRepository) > 0 {
	return u.mapper.DomainToSimpleResponseList(resultRepository)
	}
	return []dtos.UserSimpleResponse{}
}
