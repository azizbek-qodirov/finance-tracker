package kfk

import (
	pb "budget-service/genprotos"
	"encoding/json"
	"log"
)

func (kc *KafkaConsumer) handleAccountUpdate(message []byte) {
	var req pb.AccountUReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling account update message: %v", err)
		return
	}

	_, err = kc.Storage.Account().UpdateAccount(&req)
	if err != nil {
		log.Printf("Error updating account from Kafka message: %v", err)
	}
}

func (kc *KafkaConsumer) handleAccountBalanceUpdate(message []byte) {
	var req pb.AccountBalanceUReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling account balance update message: %v", err)
		return
	}

	_, err = kc.Storage.Account().UpdateBalance(&req)
	if err != nil {
		log.Printf("Error updating account balance from Kafka message: %v", err)
	}
}
