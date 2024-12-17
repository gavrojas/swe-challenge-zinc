package users

import (
	"encoding/json"
	"net/http"

	"mail_indexer_zinc/models"
	"mail_indexer_zinc/shared"
	"mail_indexer_zinc/zinc"

	"github.com/go-chi/chi/v5"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	query := `{"query": {"match_all": {}}}`
	body, err := zinc.HTTPRequestHelper("POST", "_search", []byte(query), false, false)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	body, err := zinc.HTTPRequestHelper("GET", "_doc/"+id, nil, false, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	// Recuperar el usuario autorizado desde el contexto
	user, ok := r.Context().Value(shared.AuthorizedUserKey).(map[string]interface{})
	if !ok {
		http.Error(w, "Unauthorized: User not found in context", http.StatusUnauthorized)
		return
	}

	// Retornar el usuario como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input models.UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	// Hash de la contraseña
	if err := user.HashPassword(); err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	userJSON, _ := json.Marshal(user)
	body, err := zinc.HTTPRequestHelper("POST", "_doc", userJSON, false, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := zinc.HTTPRequestHelper("DELETE", "_doc/"+id, nil, false, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userJSON, _ := json.Marshal(user)
	body, err := zinc.HTTPRequestHelper("PUT", "_doc/"+id, userJSON, false, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
