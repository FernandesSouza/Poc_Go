package dtos

import "github.com/google/uuid"


type UserSimpleResponse struct{
    Id    uuid.UUID
	Name  string
    Email string 
	Identifier string
}