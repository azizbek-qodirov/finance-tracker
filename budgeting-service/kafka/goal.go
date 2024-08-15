package kfk

import (
	pb "budget-service/genprotos"
	"encoding/json"
	"log"
)

func (kc *KafkaConsumer) handleGoalCreate(message []byte) {
	var req pb.GoalCReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling goal create message: %v", err)
		return
	}

	_, err = kc.Storage.Goal().Create(&req)
	if err != nil {
		log.Printf("Error creating goal from Kafka message: %v", err)
	}
}

func (kc *KafkaConsumer) handleGoalUpdate(message []byte) {
	var req pb.GoalUReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling goal update message: %v", err)
		return
	}

	_, err = kc.Storage.Goal().Update(&req)
	if err != nil {
		log.Printf("Error updating goal from Kafka message: %v", err)
	}
}

func (kc *KafkaConsumer) handleGoalDelete(message []byte) {
	var req pb.ByID
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Error unmarshalling goal delete message: %v", err)
		return
	}

	_, err = kc.Storage.Goal().Delete(&req)
	if err != nil {
		log.Printf("Error deleting goal from Kafka message: %v", err)
	}
}
