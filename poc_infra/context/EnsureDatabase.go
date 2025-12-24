package context

import "database/sql"

func EnsureDatabase(conn *sql.DB, dbName string) error {

	query := `
	IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = '` + dbName + `')
	BEGIN
		CREATE DATABASE [` + dbName + `]
	END
	`

	_, err := conn.Exec(query)
	return err
}
