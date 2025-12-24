package context


import (

    "database/sql"
    "log"

   _ "github.com/denisenkom/go-mssqldb"

)



func Connect() * sql.DB {

	connectionString := "sqlserver://sa:TerapiaUrgente@1309@localhost:1433?database=master"

	 db, err := sql.Open("sqlserver", connectionString)

	  if err != nil {
        log.Fatalf("Erro ao abrir conexão: %v", err)
    }

	if err = db.Ping(); err != nil {
        log.Fatalf("Erro ao conectar no SQL Server: %v", err)
    }

    return db

}