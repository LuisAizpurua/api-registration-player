package util

import (
	"encoding/json"
	"net/http"
	"time"
)

var (
	HealthHandler http.HandlerFunc
	startTime     = time.Now()
)

func init() {
	HealthHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    "UP",
			"uptime":    time.Since(startTime).String(),
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	}
}
