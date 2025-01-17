// zinc/bulk.go
package zinc

import (
	"encoding/json"
	"fmt"
	"index_data_zinc/config"
	"log"
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

	respBody, err := HTTPRequestHelper("POST", "_bulkv2", jsonData, true)
	if err != nil {
		return fmt.Errorf("error creating index: %w", err)
	}

	log.Printf("Index creation response: %s\n", respBody)
	return nil

}
