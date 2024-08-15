package kfk

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"encoding/json"
	"log"
)

func BudgetCreateHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.BudgetCReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal budget create: %v", err)
			return
		}

		_, err := db.Budget().Create(&req)
		if err != nil {
			log.Printf("Error creating budget from Kafka message: %v", err)
		}
	}
}

func BudgetUpdateHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.BudgetUReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal budget update: %v", err)
			return
		}

		_, err := db.Budget().Update(&req)
		if err != nil {
			log.Printf("Error updating budget from Kafka message: %v", err)
		}
	}
}

func BudgetDeleteHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.ByID
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal budget delete: %v", err)
			return
		}

		_, err := db.Budget().Delete(&req)
		if err != nil {
			log.Printf("Error deleting budget from Kafka message: %v", err)
		}
	}
}
