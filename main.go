package main

/*
It is the entrypoint of the application.
This class is responsible for invoking connection to database, creating server and error handler instance.
Here, the initialization of various layers take place
*/

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"takeHomeTest/controllers"
	"takeHomeTest/database"
	"takeHomeTest/errorHandlers"
	"takeHomeTest/repository"
	"takeHomeTest/routes"
	"takeHomeTest/service"
)

var (
	dbInstance   *sql.DB
	repo         repository.Repository
	serv         service.Service
	controller   controllers.Controller
	errorHandler errorHandlers.ErrorHandlers
)

func main() {
	log.Info("Hi, this is Vestiaire Collective take home test")
	initDatabase()
	createServer()
}

func createServer() {

	server := gin.Default()
	initializeLayers()
	routes.RegisterHandlers(server, controller).RegisterHandlers()

	err := server.Run()

	if err != nil {
		errorHandler.FailOnError(err, "Unable to start server")
		dbInstance.Close()
	}
}

func initializeLayers() {
	repo = repository.NewRepository(dbInstance, errorHandler)
	serv = service.NewService(repo, errorHandler)
	controller = controllers.NewController(serv, errorHandler)
}

func initDatabase() {
	log.Info("Connecting to MySql...")
	createErrorHandler()
	dbInstance = database.GetDbConnection()
}

func createErrorHandler() {
	errorHandler = errorHandlers.NewErrorHandler()
}
