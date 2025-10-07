package infrastructure

import "dia1/orders/domain"

type InMemoryOrderRepository struct {
	data map[string]*domain.Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		data: make(map[string]*domain.Order),
	}
}

func (r *InMemoryOrderRepository) Save(order *domain.Order) error {
	r.data[order.ID] = order
	return nil

}
