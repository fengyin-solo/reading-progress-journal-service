package handler

import (
	"encoding/json"
	"net/http"

	"readinglog/internal/service"
)

func SaveRetryScenario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(service.RunSaveRetryScenario("entry-42"))
}
