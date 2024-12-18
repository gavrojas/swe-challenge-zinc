package emails

import (
	"encoding/json"
	"mail_indexer_zinc/models"
	"mail_indexer_zinc/zinc"
	"net/http"
)

func SearchMailsByField(w http.ResponseWriter, r *http.Request) {
	var req models.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	searchQuery := map[string]interface{}{
		"search_type": "match",
		"query": map[string]interface{}{
			"term":  req.Term,
			"field": req.Field,
		},
		"from":        req.From,
		"max_results": req.MaxResults,
	}

	searchBody, err := json.Marshal(searchQuery)
	if err != nil {
		http.Error(w, `{"error": "Failed to prepare search request"}`, http.StatusInternalServerError)
		return
	}

	respBody, err := zinc.HTTPRequestHelper("POST", "_search", searchBody, false, true)
	if err != nil {
		http.Error(w, `{"error": "Failed to contact ZincSearch"}`, http.StatusInternalServerError)
		return
	}

	// Establecer el tipo de contenido de la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
