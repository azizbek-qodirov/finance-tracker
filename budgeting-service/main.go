package main

import (
	"log"
	"net"

	cf "budget-service/config"
	kfk "budget-service/kafka"
	"budget-service/service"
	"budget-service/storage"
)

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()

	db, err := storage.NewDatabaseStorage(config)
	em.CheckErr(err)
	defer db.Close()

	listener, err := net.Listen("tcp", config.BUDGET_SERVICE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := service.InitServer(db)
	kfk.InitKafka(&config, db)

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
