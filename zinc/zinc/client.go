// zinc/client.go
package zinc

import (
	"bytes"
	"fmt"
	"index_data_zinc/config"
	"io"
	"net/http"
)

func HTTPRequestHelper(method, path string, body []byte, create bool) ([]byte, error) {
	var url string
	if create {
		url = fmt.Sprintf("%s/%s", config.ZincHost, path)
	} else {
		url = fmt.Sprintf("%s/%s/%s", config.ZincHost, config.ZincIndex, path)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.ZincUsername, config.ZincPassword)
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

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("ZincSearch returned error: %s", resp.Status)
	}
	return respBody, nil
}
