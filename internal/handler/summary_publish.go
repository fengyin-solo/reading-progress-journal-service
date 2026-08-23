package handler

import (
	"net/http"

	"readinglog/internal/service"
	"readinglog/pkg/httpx"
)

func SummaryPublishScenario(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, service.RunSummaryPublishScenario())
}
