package services

import (
	dtos "poc_api/poc_application/dataTransferObjects/userDtos/request"
	"poc_api/poc_application/interfaces"
	"poc_api/poc_infra/interfaces"
	"github.com/google/uuid"
)

type UserCommandService struct {
	repository infra_interfaces.IUserRepository
	mapper     application_interfaces.IUserMapper
}


func NewUserCommandService(
	userRepository infra_interfaces.IUserRepository,
	userMapper application_interfaces.IUserMapper,
) application_interfaces.IUserCommandService {
	return &UserCommandService{
		repository: userRepository,
		mapper:     userMapper,
	}
}

func (s *UserCommandService) Register(dto *dtos.UserRegisterRequest) string {

	user := s.mapper.DtoRegisterToDomain(dto)

	for {
		newId := uuid.New()

		exists := s.repository.CheckExistId(newId)
		if !exists {
			user.Id = newId
			break
		}
	}


	var result = s.repository.Save(&user)

	if(result != nil){
		return "Erro ao realizar cadastro"
	}

	return "Sucesso !"
}


