// zinc/bulk.go
package zinc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"index_data_zinc/config"
	"io"
	"net/http"
)

func SendBulkToZinc(documents []EmailDocument) error {
	bulkRequest := BulkV2Request{
		Index:   config.ZincIndex,
		Records: documents,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(bulkRequest)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/_bulkv2", config.ZincHost), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.SetBasicAuth(config.ZincUsername, config.ZincPassword)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check response
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("error response from Zinc: %s - %s", resp.Status, string(body))
	}
	return nil
}
