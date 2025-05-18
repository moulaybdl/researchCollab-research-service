package main

import (
	"fmt"
	"log/slog"
	"moulaybdl/researchCollab/researchSevice/internal/adapters/repository"
	"moulaybdl/researchCollab/researchSevice/internal/logger"
	"net/http"
	"os"
)


func main() {
	// define here customized ports, DSN
	port := "8080"

	// define here the DSN


	logger.Init()
	logger.Logger.Info("Starting Now ...")


	// define here any services (db, cache, mail server, ...)

	// open a connection to the database
  _, err := repository.OpenDB("DSN")
  if err != nil {
    logger.Logger.Error(err.Error())
    return
  }
  logger.Logger.Info("Connection to database established !")



	// Server setup:
	srv := &http.Server{
		Addr: port,
		Handler: InitRoutes(),
		ErrorLog: slog.NewLogLogger(logger.Logger.Handler(), slog.LevelInfo),
	}

	err = srv.ListenAndServe()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("Error starting the server: %v", err))
	}	
		os.Exit(0)
}