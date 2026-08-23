package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"readinglog/internal/handler"
	"readinglog/internal/model"
)

func TestSummaryPublishRetryIsIdempotentAndStaleCallbackCannotRegressState(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler.SummaryPublishScenario(recorder, httptest.NewRequest(http.MethodPost, "/summary-publish", nil))
	if recorder.Code != http.StatusOK {
		t.Fatalf("summary publish returned %d: %s", recorder.Code, recorder.Body.String())
	}
	var response struct {
		Data model.SummaryPublishResult `json:"data"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode summary publish response: %v", err)
	}
	if response.Data.Attempts != 2 {
		t.Errorf("expected one retry, attempts=%d", response.Data.Attempts)
	}
	if response.Data.Deliveries != 1 {
		t.Errorf("retry duplicated summary delivery: %d", response.Data.Deliveries)
	}
	if response.Data.StoreState != model.PublishCompleted || response.Data.StoreVersion != 2 {
		t.Errorf("stale callback regressed stored job to %s@%d", response.Data.StoreState, response.Data.StoreVersion)
	}
	if response.Data.CacheState != model.PublishCompleted || response.Data.CacheVersion != 2 {
		t.Errorf("stale callback regressed cached job to %s@%d", response.Data.CacheState, response.Data.CacheVersion)
	}
}
