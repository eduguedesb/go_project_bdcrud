package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Criar um usuário
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := db.Exec("INSERT INTO users (nome, email) VALUES (?, ?)", user.Nome, user.Email)
	if err != nil {
		http.Error(w, "Erro ao inserir usuário", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Usuário criado com sucesso!")
}

// Ler todos os usuários
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, nome, email FROM users")
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Nome, &user.Email)
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// Ler um usuário específico
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user User
	err := db.QueryRow("SELECT id, nome, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Nome, &user.Email)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Atualizar um usuário
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := db.Exec("UPDATE users SET nome = ?, email = ? WHERE id = ?", user.Nome, user.Email, id)
	if err != nil {
		http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Usuário atualizado com sucesso!")
}

// Deletar um usuário
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Erro ao deletar usuário", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Usuário deletado com sucesso!")
}
