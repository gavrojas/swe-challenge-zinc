package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"mail_indexer_zinc/zinc"
	"net/http"
	"strconv"
	"time"
)

type ContextKey string

const AuthorizedUserKey ContextKey = "authorizedUser"

func AuthenticateSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenStr := GetTokenFromRequest(r)
		claims, err := ParseJWTToken(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		sessionData, sessionExists := Sessions[claims.Session]
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
				"term": {
					"id": "%s"
				}
			}
		}`, strconv.FormatUint(uint64(sessionData.Uid), 10))

		response, err := zinc.HTTPRequestHelper("POST", "_search", []byte(query), false)
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
