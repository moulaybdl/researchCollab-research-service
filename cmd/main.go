package main

import (
	"fmt"
	"log"
	"moulaybdl/researchCollab/researchSevice/internal/config"
	"moulaybdl/researchCollab/researchSevice/internal/logger"
	"net"
	"os"

	pb "moulaybdl/researchCollab/researchSevice/internal/core/ports/protobufs/protobufs"
	"moulaybdl/researchCollab/researchSevice/internal/core/services"

	"google.golang.org/grpc"
)

// info:
// database name: researchcollab_research



func main() {

	// Init the logger
	logger.Init()
	logger.Logger.Info("Starting Now ...")
	
	// Load configuration
	logger.Logger.Info("Loading configuration...")
	cfg := config.Config{}
	err := config.LoadConfig(&cfg)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("Error loading configuration: %v", err))
		os.Exit(1)
	}
	logger.Logger.Info("Configuration Loaded !")


	// define here any services (db, cache, mail server, ...)

	// open a connection to the database
	// logger.Logger.Info("Connecting to the database...")
  // _, err = repository.OpenDB(cfg.DSN)
  // if err != nil {
  //   logger.Logger.Error(err.Error())
  //   return
  // }
  // logger.Logger.Info("Connection to database established !")


	// logger.Logger.Info("Starting the server...")
	// // Server setup:
	// srv := &http.Server{
	// 	Addr: fmt.Sprintf(":%s", cfg.Port),
	// 	Handler: InitRoutes(),
	// 	ErrorLog: slog.NewLogLogger(logger.Logger.Handler(), slog.LevelInfo),
	// }

	// err = srv.ListenAndServe()
	// if err != nil {
	// 	logger.Logger.Error(fmt.Sprintf("Error starting the server: %v", err))
	// }	
	// 	os.Exit(0)

	// ###################### gRPC THINGS ################################

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		}

		// Create gRPC server
		s := grpc.NewServer()
		
		// Register service
		pb.RegisterGreeterServer(s, &services.Server{})
		
		log.Println("gRPC server listening on port 50051")
		
		// Start serving
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
	}


}

// commands for you
// migrate create -seq -ext=.sql -dir=./migrations create_movies_table
// migrate -path=./migrations -database= up