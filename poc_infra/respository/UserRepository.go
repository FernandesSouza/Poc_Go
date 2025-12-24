package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"poc_api/poc_domain/entities"
	"poc_api/poc_infra/interfaces"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) infra_interfaces.IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Save(user *entities.User) error {
	_, err := u.db.Exec(
		"INSERT INTO dbo.Users (Id, Name, Email, Identifier) VALUES (@p1, @p2, @p3, @p4)",
		user.Id,
		user.Name,
		user.Email,
		user.Identifier,
	)

	return err
}

func (r *UserRepository) FindLast() *entities.User {
	query := `
        SELECT TOP 1 Id, Name, Email, Identifier
        FROM dbo.Users
        ORDER BY Id DESC;`

	row := r.db.QueryRow(query)

	var u entities.User

	err := row.Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.Identifier,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return nil
	}

	return &u
}

func (u *UserRepository) FindAllUsers() []entities.User {

	query := `
		SELECT Id, Name, Email, Identifier
		FROM dbo.Users
		ORDER BY Id DESC;
	`

	rows, err := u.db.Query(query)

	if err != nil {
		return []entities.User{}
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {

		var user entities.User

		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Identifier,
		); err != nil {
			continue
		}

		users = append(users, user)
	}

	return users
}

func (u *UserRepository) CheckExistId(id uuid.UUID) bool {
	var exists int

	query := `
		SELECT COUNT(1)
		FROM dbo.Users
		WHERE Id = @p1
	`

	err := u.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false
	}

	return exists > 0
}


func (u *UserRepository) FindByEmail(email string) entities.User{
	query := `
	SELECT * FROM 
	dbo.Users
	WHERE Email = @p1
	`

	row := u.db.QueryRow(query)

	var user entities.User

	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Identifier,
	)

	if err != nil {
		return entities.User{}
	}
		
	return user	
}
