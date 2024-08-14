package handlers

import (
	pb "gateway-service/genprotos"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Lesson pb.LessonServiceClient
}

func NewHandler(connL *grpc.ClientConn) *HTTPHandler {
	return &HTTPHandler{
		Lesson: pb.NewLessonServiceClient(connL),
	}
}
