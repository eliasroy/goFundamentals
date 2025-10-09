package handler

import (
	"dia3/model"
	"dia3/service"
	"encoding/json"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	// Canales para recibir los resultados
	ch := make(chan []model.User)

	//lanza ambas tareas concurrentemente
	go func() {
		ch <- service.FetchUsersFromSourceA()
	}()
	go func() {
		ch <- service.FetchUsersFromSourceB()
	}()

	var users []model.User

	// Espera a recibir los resultados de ambas tareas
	for i := 0; i < 2; i++ {
		u := <-ch
		users = append(users, u...)
	}

	// Envía la respuesta JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
