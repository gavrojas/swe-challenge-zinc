// config/config.go
package config

import "os"

var (
	ZincHost     = os.Getenv("ZINC_URL")
	ZincIndex    = "enron_emails_test"
	ZincUsername = os.Getenv("ZINC_USERNAME")
	ZincPassword = os.Getenv("ZINC_PASSWORD")
	BatchSize    = 10
	NumWorkers   = 1
	// NumWorkers   = 4 -> más recursos
)
