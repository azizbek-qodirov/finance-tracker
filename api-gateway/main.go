package main

import (
	"fmt"
	"gateway-service/api"
	cf "gateway-service/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()

	BudgetConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", config.BUDGET_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer BudgetConn.Close()

	r := api.NewRouter(BudgetConn)

	fmt.Printf("Server started on port %s\n", config.HTTP_PORT)
	if err := r.Run(config.HTTP_PORT); err != nil {
		panic(err)
	}
}
