package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"readinglog/internal/handler"
	"readinglog/internal/model"
)

func TestExcerptBatchOwnershipSurvivesBufferReuseAndDelayedConsumers(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler.SliceScenario(recorder, httptest.NewRequest(http.MethodPost, "/slice-scenario", nil))
	if recorder.Code != http.StatusOK {
		t.Fatalf("slice scenario returned %d: %s", recorder.Code, recorder.Body.String())
	}
	var result model.SliceScenarioResult
	if err := json.Unmarshal(recorder.Body.Bytes(), &result); err != nil {
		t.Fatalf("decode slice scenario: %v", err)
	}
	if result.ParsedFirst != "alpha" {
		t.Errorf("first parsed batch was overwritten by buffer reuse: %q", result.ParsedFirst)
	}
	if result.DelayedReceipt != "cedar" {
		t.Errorf("delayed receipt observed caller mutation: %q", result.DelayedReceipt)
	}
	if result.CachedReplay != "delta" {
		t.Errorf("cached batch was mutated through returned slice: %q", result.CachedReplay)
	}
	if result.Exported != "echo" {
		t.Errorf("queued export observed producer buffer reuse: %q", result.Exported)
	}
}
