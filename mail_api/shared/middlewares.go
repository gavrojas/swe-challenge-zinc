package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"mail_indexer_zinc/zinc"
	"net/http"
	"time"
)

type ContextKey string

const AuthorizedUserKey ContextKey = "authorizedUser"

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := []string{
			"http://localhost:5173",
			"https://swe-challenge-zinc-gavrojas.vercel.app",
		}

		origin := r.Header.Get("Origin")
		for _, o := range allowedOrigins {
			if o == origin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthenticateSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenStr := GetTokenFromRequest(r)
		fmt.Println("token", tokenStr)
		claims, err := ParseJWTToken(tokenStr)
		fmt.Println("err", err)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		sessionData, sessionExists := Sessions[claims.Session]

		fmt.Println("claims", claims)
		fmt.Println("claims session", claims.Session)
		fmt.Println("session data", sessionData)
		fmt.Println("session sessionExists", sessionData)

		if !sessionExists {
			http.Error(w, "Unauthorized: Session not found", http.StatusUnauthorized)
			return
		}

		if sessionData.ExpiryTime.Before(time.Now()) {
			http.Error(w, "Unauthorized: Session expired", http.StatusUnauthorized)
			return
		}

		query := fmt.Sprintf(`{
			"query": {
				"match": {
					"id": "%s"
				}
			}
		}`, sessionData.Uid)

		fmt.Println("sessionData.Uid", sessionData.Uid)
		fmt.Println("query", query)

		response, err := zinc.HTTPRequestHelper("POST", "_search", []byte(query), false, false)

		if err != nil {
			http.Error(w, "Error fetching user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var searchResult map[string]interface{}
		if err := json.Unmarshal(response, &searchResult); err != nil {
			http.Error(w, "Error parsing ZincSearch response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Validar si el usuario existe
		hits := searchResult["hits"].(map[string]interface{})["hits"].([]interface{})
		if len(hits) == 0 {
			http.Error(w, "Unauthorized: User not found", http.StatusUnauthorized)
			return
		}

		user := hits[0].(map[string]interface{})["_source"]

		ctx := context.WithValue(r.Context(), AuthorizedUserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
