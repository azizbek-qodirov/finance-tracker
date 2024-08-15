package kfk

import (
	pb "budget-service/genprotos"
	"encoding/json"
	"log"
)

func (kc *KafkaConsumer) handleTransactionCreate(message []byte) {
	var req pb.TransactionCReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling transaction create message: %v", err)
		return
	}

	_, err = kc.Storage.Transaction().Create(&req)
	if err != nil {
		log.Printf("Error creating transaction from Kafka message: %v", err)
	}
}

func (kc *KafkaConsumer) handleTransactionDelete(message []byte) {
	var req pb.ByID
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling transaction delete message: %v", err)
		return
	}

	_, err = kc.Storage.Transaction().Delete(&req)
	if err != nil {
		log.Printf("Error deleting transaction from Kafka message: %v", err)
	}
}
