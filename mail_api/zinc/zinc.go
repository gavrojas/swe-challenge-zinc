package zinc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mail_indexer_zinc/models"
	"net/http"
)

const (
	ZINC_URL         = "http://zinc:4080/api"
	ZINC_INDEX_USERS = "users"
	ZINC_INDEX_EMAIL = "email"
	ZINC_USERNAME    = "admin"
	ZINC_PASSWORD    = "Complexpass#123"
)

func IndexExists(indexName string) (bool, error) {
	existsQuery := fmt.Sprintf(`{
		"query": {
			"match": {
				"username": "%s"
			}
		}
	}`, ZINC_USERNAME)

	// Hacer la solicitud HTTP para verificar la existencia del índice
	respBody, err := HTTPRequestHelper("POST", "_search", []byte(existsQuery), false)
	if err != nil {
		return false, nil
	}

	// Analizar la respuesta para determinar si el índice existe
	var existsResponse models.ExistsIndexResponse

	if err := json.Unmarshal(respBody, &existsResponse); err != nil {
		return false, fmt.Errorf("error parsing existence check response: %w", err)
	}
	fmt.Println("response index exist: ", existsResponse.Hits.Total.Value)
	// Si el índice ya existe, retornar true
	if existsResponse.Hits.Total.Value >= 0 {
		fmt.Printf("Index '%s' already exists.\n", indexName)
		return true, nil
	}
	// Si el índice no existe, retornar false
	return false, nil
}

// Generic HTTP request to ZincSearch
func HTTPRequestHelper(method, path string, body []byte, create bool) ([]byte, error) {
	var url string
	if create {
		url = fmt.Sprintf("%s/%s", ZINC_URL, path)
	} else {
		url = fmt.Sprintf("%s/%s/%s", ZINC_URL, ZINC_INDEX_USERS, path)
	}

	fmt.Println("url", url)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(ZINC_USERNAME, ZINC_PASSWORD)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta para diagnóstico
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Imprimir detalles de la respuesta para depuración
	fmt.Printf("Zinc response status: %s\n", resp.Status)
	fmt.Printf("Zinc response body: %s\n", string(respBody))

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("ZincSearch returned error: %s", resp.Status)
	}
	return respBody, nil
}

func CreateIndex(name string) error {
	exists, err := IndexExists(name)
	if err != nil {
		return fmt.Errorf("error checking if index exists: %w", err)
	}

	if exists {
		return nil
	}
	indexBody := fmt.Sprintf(`{
		"name": "%s",
		"mappings": {
			"properties": {
				"id": {
					"type": "keyword"
				},
				"username": {
					"type": "keyword"
				},
				"password": {
					"type": "text"
				}
			}
		}
	}`, name)

	respBody, err := HTTPRequestHelper("POST", "index", []byte(indexBody), true)
	if err != nil {
		return fmt.Errorf("error creating index: %w", err)
	}

	fmt.Printf("Index creation response: %s\n", respBody)
	return nil
}
