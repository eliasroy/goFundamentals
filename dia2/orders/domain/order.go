package domain

import "errors"

type OrderItem struct {
	Product  string
	Quantity int
	Price    float64
}

type Order struct {
	ID     string
	Items  []OrderItem
	Status string
}

// comportamiento del dominio
func (o *Order) calculateTotal() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += float64(item.Quantity) * item.Price
	}
	return total
}

func (e *Order) Validate() error {
	if len(e.Items) == 0 {
		return errors.New("an order must have at least one item")
	}

	if e.Status == "" {
		return errors.New("an order must have a status")
	}

	return nil
}
