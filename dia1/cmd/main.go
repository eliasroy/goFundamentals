package cmd

import (
	"dia1/orders/application"
	"dia1/orders/infrastructure"
	"dia1/orders/interfaces"
	"fmt"
	"net/http"
)

func main() {
	repo := infrastructure.NewInMemoryOrderRepository()
	service := application.NewOrderService(repo)
	handler := interfaces.NewOrderHandler(service)

	http.HandleFunc("/orders", handler.CreateOrder)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
