package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Conectar ao banco de dados
	connectDB()

	// Configurar as rotas
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Iniciar o servidor
	log.Println("Servidor iniciado na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
