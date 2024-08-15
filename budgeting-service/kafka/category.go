package kfk

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"encoding/json"
	"log"
)

func CategoryCreateHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.CategoryCReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal category create: %v", err)
			return
		}

		_, err := db.Category().Create(&req)
		if err != nil {
			log.Printf("Error creating category from Kafka message: %v", err)
		}
	}
}

func CategoryUpdateHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.CategoryUReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal category update: %v", err)
			return
		}

		_, err := db.Category().Update(&req)
		if err != nil {
			log.Printf("Error updating category from Kafka message: %v", err)
		}
	}
}

func CategoryDeleteHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.ByID
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal category delete: %v", err)
			return
		}

		_, err := db.Category().Delete(&req)
		if err != nil {
			log.Printf("Error deleting category from Kafka message: %v", err)
		}
	}
}
