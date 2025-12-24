package infra_interfaces

import (
    
    "poc_api/poc_domain/entities"
    "github.com/google/uuid"
)

type IUserRepository interface {
    FindLast() *entities.User
    FindByEmail(email string) entities.User
    CheckExistId(id uuid.UUID) bool
    FindAllUsers() []entities.User
    Save(user *entities.User) error
}
