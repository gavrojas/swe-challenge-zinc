package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"mail_indexer_zinc/models"
	"mail_indexer_zinc/shared"
	"mail_indexer_zinc/zinc"

	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.UserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"username": "%s"
			}
		}
	}`, userInput.Username)

	fmt.Println("query: ", query)

	body, err := zinc.HTTPRequestHelper("POST", "_search", []byte(query), false, false)
	fmt.Println("body: ", body)
	fmt.Println("Query bytes: ", []byte(query))

	if err != nil {
		http.Error(w, "Error checking if user exists", http.StatusInternalServerError)
		return
	}

	// Si ya existe, devolver un conflicto
	var searchResponse models.SearchUserResponse

	if err := json.Unmarshal(body, &searchResponse); err != nil {
		http.Error(w, "Error parsing search response", http.StatusInternalServerError)
		return
	}

	fmt.Println("response register:", searchResponse)

	for _, hit := range searchResponse.Hits.Hits {
		if hit.Source.Username == userInput.Username {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := map[string]interface{}{
		"username": userInput.Username,
	}

	user["password"] = string(hashedPassword)
	userJSON, _ := json.Marshal(user)

	// Almacenar el usuario en ZincSearch
	responseBody, err := zinc.HTTPRequestHelper("POST", "_doc", userJSON, false, false)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	var createResponse struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(responseBody, &createResponse); err != nil {
		http.Error(w, "Error parsing create response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"username": userInput.Username,
		"id":       createResponse.ID,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.UserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"username": "%s"
			}
		}
	}`, userInput.Username)

	body, err := zinc.HTTPRequestHelper("POST", "_search", []byte(query), false, false)
	if err != nil {
		http.Error(w, "Error checking user credentials", http.StatusInternalServerError)
		return
	}

	var searchResponse models.SearchUserResponse

	if err := json.Unmarshal(body, &searchResponse); err != nil {
		http.Error(w, "Error parsing search response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("response Login", searchResponse)

	var userFound bool
	var user models.User
	for _, hit := range searchResponse.Hits.Hits {
		if hit.Source.Username == userInput.Username {
			user = models.User{
				ID:       hit.ID,
				Password: hit.Source.Password,
			}
			userFound = true
			break
		}
	}
	if !userFound {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Comparar la contraseña hasheada
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewV5(uuid.UUID{}, "session").String()

	session := shared.Session{
		Uid: user.ID,
		// 24 horas
		ExpiryTime: time.Now().Add(time.Hour * 24),
	}

	shared.Sessions[sessionToken] = session

	claims := shared.Payload{
		MapClaims: jwt.MapClaims{
			"iat": jwt.NewNumericDate(time.Now()),                       // issued at,
			"eat": jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), // expired at - tiempo que va a durar el token
		},
		Session: sessionToken,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": tokenString,
	})

}

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenStr := shared.GetTokenFromRequest(r)

	if _, ok := shared.Sessions[tokenStr]; !ok {
		delete(shared.Sessions, tokenStr)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged out successfully"))
		return
	}

	http.Error(w, "Invalid token", http.StatusUnauthorized)
}
