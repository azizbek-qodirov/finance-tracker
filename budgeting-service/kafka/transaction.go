package kfk

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"encoding/json"
	"log"
)

func TransactionCreateHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.TransactionCReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal transaction create: %v", err)
			return
		}

		_, err := db.Transaction().Create(&req)
		if err != nil {
			log.Printf("Error creating transaction from Kafka message: %v", err)
		}
	}
}

func TransactionDeleteHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.ByID
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal transaction delete: %v", err)
			return
		}

		_, err := db.Transaction().Delete(&req)
		if err != nil {
			log.Printf("Error deleting transaction from Kafka message: %v", err)
		}
	}
}
