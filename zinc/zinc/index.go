package zinc

import (
	"fmt"
	"index_data_zinc/config"
	"io"
	"log"
	"net/http"
)

func CreateZincIndex(name string) error {
	exists, err := IndexExists(name)
	if err != nil {
		return fmt.Errorf("error checking if index exists: %w", err)
	}

	if exists {
		log.Printf("Index %s already exists", name)
		return nil
	}

	indexBody := fmt.Sprintf(`{
		"name": "%s",
		"storage_type": "disk",
		"mappings": {
			"properties": {
				"message_id": {"type": "keyword"},
				"date": {"type": "text"},
				"from": {"type": "text"},
				"to": {"type": "text"},
				"subject": {"type": "text"},
				"body": {"type": "text"},
				"mime_version": {"type": "text"},
				"content_type": {"type": "text"},
				"content_transfer_encoding": {"type": "text"},
				"x_from": {"type": "text"},
				"x_to": {"type": "text"},
				"x_cc": {"type": "text"},
				"x_bcc": {"type": "text"},
				"x_folder": {"type": "text"},
				"x_origin": {"type": "text"},
				"x_filename": {"type": "text"},
				"user": {"type": "keyword"},
				"folder_path": {"type": "keyword"}
			}
		}
	}`, name)

	respBody, err := HTTPRequestHelper("POST", "index", []byte(indexBody), true)
	if err != nil {
		return fmt.Errorf("error creating index: %w", err)
	}

	log.Printf("Index creation response: %s\n", respBody)
	return nil
}

func IndexExists(name string) (bool, error) {
	url := fmt.Sprintf("%s/index/%s", config.ZincHost, name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.SetBasicAuth(config.ZincUsername, config.ZincPassword)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true, nil
	}
	if resp.StatusCode == 404 {
		return false, nil
	}

	respBody, _ := io.ReadAll(resp.Body)
	return false, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
}
