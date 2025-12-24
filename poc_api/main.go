package main

import (
	"database/sql"
	"log"
	controller "poc_api/poc_api/controllers"
	route "poc_api/poc_api/routes"
	"poc_api/poc_application/mappers"
	"poc_api/poc_application/services"
	repository "poc_api/poc_infra/respository"
	database "poc_api/poc_infra/context"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

func main() {

	masterConnString := "server=localhost;user id=sa;password=TerapiaUrgente!1309;encrypt=disable"
   appConnString := "server=localhost;user id=sa;password=TerapiaUrgente!1309;database=MyDatabase;encrypt=disable"


	masterDb, err := sql.Open("sqlserver", masterConnString)
	if err != nil {
		log.Fatal(err)
	}

	if err = masterDb.Ping(); err != nil {
		log.Fatal(err)
	}


	err = database.EnsureDatabase(masterDb, "MyDatabase")
	if err != nil {
		log.Fatal(err)
	}

	masterDb.Close()

	db, err := sql.Open("sqlserver", appConnString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	err = database.EnsureUsersTable(db)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userMapper := mappers.NewUserMapper()
	userCommandService := services.NewUserCommandService(userRepo, userMapper)
	userQueryService := services.NewUserQueryService(userRepo, userMapper)

	userController := controller.NewUserController(userCommandService, userQueryService)

	r := gin.Default()
	route.RegisterUserRoutes(r, userController)
	r.Run(":8080")
}
