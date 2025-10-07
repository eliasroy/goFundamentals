package application

import (
	"dia1/orders/domain"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Save(order *domain.Order) error
}
type OrderService struct {
	repo OrderRepository
}

func NewOrderService(r OrderRepository) *OrderService {
	return &OrderService{
		repo: r,
	}
}

func (s *OrderService) CreateOrder(items []domain.OrderItem) (*domain.Order, error) {

	order := &domain.Order{
		ID:     uuid.New().String(),
		Items:  items,
		Status: "created",
	}

	if err := order.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Save(order); err != nil {
		return nil, err
	}

	return order, nil
}
