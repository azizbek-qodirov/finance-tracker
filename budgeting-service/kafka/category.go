package kfk

import (
	pb "budget-service/genprotos"
	"encoding/json"
	"log"
)

func (kc *KafkaConsumer) handleCategoryCreate(message []byte) {
	var req pb.CategoryCReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling category create message: %v", err)
		return
	}

	_, err = kc.Storage.Category().Create(&req) // Assuming your storage has this method
	if err != nil {
		log.Printf("Error creating category from Kafka message: %v", err)
	}
}

func (kc *KafkaConsumer) handleCategoryUpdate(message []byte) {
	var req pb.CategoryUReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling category update message: %v", err)
		return
	}

	_, err = kc.Storage.Category().Update(&req)
	if err != nil {
		log.Printf("Error updating category from Kafka message: %v", err)
	}
}

func (kc *KafkaConsumer) handleCategoryDelete(message []byte) {
	var req pb.ByID
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling category delete message: %v", err)
		return
	}

	_, err = kc.Storage.Category().Delete(&req)
	if err != nil {
		log.Printf("Error deleting category from Kafka message: %v", err)
	}
}
