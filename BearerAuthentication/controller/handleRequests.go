package controller

import (
	"encoding/json"
	"net/http"
	"bearerAuth/repository"
	"bearerAuth/model"
	"io"
	"fmt"
	"github.com/google/uuid"
)

var userRepo = repository.NewUserRepository() 
var tokenStore = make(map[string]int)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userRepo.Users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds model.LoginDTO
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := json.Unmarshal(body, &creds); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received User: %+v\n", creds)


	email, password :=  creds.Email, creds.Password
	fmt.Println(email, password)
	if email == "" || password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	u, err := userRepo.FindByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if u.Authenticate(email, password) {
		w.WriteHeader(http.StatusOK)
		newUUID, _ := uuid.NewUUID()
		tokenStore["Bearer " + newUUID.String()] = u.Id
		w.Write([]byte("Bearer " + newUUID.String()))
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}


func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := r.Header.Get("Authorization")
		_, exists := tokenStore[token]
		if !exists {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}