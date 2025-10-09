package service

import "time"
import "dia3/model"

// Simula obtener usuarios de una fuente A (tarda 1 segundo)
func FetchUsersFromSourceA() [] model.User {
	time.Sleep(1 * time.Second)
	return []model.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com"}
	}
}

func FetchUsersFromSourceB() []model.User {
	time.Sleep(2 * time.Second)
	return []model.User{
		{ID: 4, Name: "David", Email: "david@example.com"},
		{ID: 5, Name: "Eve", Email: "eve@example.com"},
		{ID: 6, Name: "Frank", Email: "frank@example.com"}
	}
}