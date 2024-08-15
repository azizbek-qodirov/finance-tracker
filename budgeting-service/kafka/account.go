package kfk

import (
	"encoding/json"
	"log"

	pb "budget-service/genprotos"
	"budget-service/storage"
)

func AccountUpdateHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.AccountUReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal account update: %v", err)
			return
		}

		_, err := db.Account().UpdateAccount(&req)
		if err != nil {
			log.Printf("Error updating account from Kafka message: %v", err)
		}
	}
}

func AccountBalanceUpdateHandler(db *storage.Storage) func(message []byte) {
	return func(message []byte) {
		var req pb.AccountBalanceUReq
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("Failed to unmarshal account balance update: %v", err)
			return
		}

		_, err := db.Account().UpdateBalance(&req)
		if err != nil {
			log.Printf("Error updating account balance from Kafka message: %v", err)
		}
	}
}
