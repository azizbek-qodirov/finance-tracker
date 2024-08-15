package kfk

import (
	pb "budget-service/genprotos"
	"encoding/json"
	"log"
)

func (kc *KafkaConsumer) handleBudgetCreate(message []byte) {
	var req pb.BudgetCReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling budget create message: %v", err)
		return
	}

	_, err = kc.Storage.Budget().Create(&req)
	if err != nil {
		log.Printf("Error creating budget from Kafka message: %v", err)
	}
}

func (kc *KafkaConsumer) handleBudgetUpdate(message []byte) {
	var req pb.BudgetUReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling budget update message: %v", err)
		return
	}

	_, err = kc.Storage.Budget().Update(&req)
	if err != nil {
		log.Printf("Error updating budget from Kafka message: %v", err)
	}
}

func (kc *KafkaConsumer) handleBudgetDelete(message []byte) {
	var req pb.ByID
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling budget delete message: %v", err)
		return
	}

	_, err = kc.Storage.Budget().Delete(&req)
	if err != nil {
		log.Printf("Error deleting budget from Kafka message: %v", err)
	}
}
