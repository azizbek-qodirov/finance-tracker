package kfk

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"encoding/json"
	"log"
)

func GoalCreateHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.GoalCReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal goal create: %v", err)
			return
		}

		_, err := db.Goal().Create(&req)
		if err != nil {
			log.Printf("Error creating goal from Kafka message: %v", err)
		}
	}
}

func GoalUpdateHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.GoalUReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal goal update: %v", err)
			return
		}

		_, err := db.Goal().Update(&req)
		if err != nil {
			log.Printf("Error updating goal from Kafka message: %v", err)
		}
	}
}

func GoalDeleteHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.ByID
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal goal delete: %v", err)
			return
		}

		_, err := db.Goal().Delete(&req)
		if err != nil {
			log.Printf("Error deleting goal from Kafka message: %v", err)
		}
	}
}
