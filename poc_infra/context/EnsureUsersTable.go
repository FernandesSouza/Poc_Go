package context

import "database/sql"

func EnsureUsersTable(conn *sql.DB) error {

	query := `
	IF NOT EXISTS (
		SELECT * FROM sys.tables WHERE name = 'Users'
	)
	BEGIN
		CREATE TABLE Users (
			Id UNIQUEIDENTIFIER NOT NULL,
			Name NVARCHAR(200) NOT NULL,
			Email NVARCHAR(200) NOT NULL,
			Identifier NVARCHAR(50) NOT NULL,

			CONSTRAINT PK_Users PRIMARY KEY (Id),
		)
	END
	`

	_, err := conn.Exec(query)
	return err
}
