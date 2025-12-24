package application_interfaces

import(
	dtos "poc_api/poc_application/dataTransferObjects/userDtos/response"
)


type IUserQueryService interface{
	FindAll() []dtos.UserSimpleResponse
}