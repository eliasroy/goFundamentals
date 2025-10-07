package interfaces

import (
	"dia1/orders/application"
	"dia1/orders/domain"
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	service *application.OrderService
}

func NewOrderHandler(s *application.OrderService) *OrderHandler {
	return &OrderHandler{
		service: s,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var domainItems []domain.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&domainItems); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := h.service.CreateOrder(domainItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
