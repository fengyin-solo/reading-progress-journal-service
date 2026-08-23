package handler

import (
	"encoding/json"
	"net/http"

	"readinglog/internal/service"
)

func SliceScenario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(service.RunSliceScenario())
}
