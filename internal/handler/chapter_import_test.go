package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"readinglog/internal/handler"
	"readinglog/internal/model"
)

func TestChapterImportFailureClosesReadersAndPreservesError(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler.ChapterImportScenario(recorder, httptest.NewRequest(http.MethodPost, "/chapter-import", nil))
	if recorder.Code < http.StatusBadRequest {
		t.Fatalf("chapter import returned %d: %s", recorder.Code, recorder.Body.String())
	}
	var response struct {
		Code int                       `json:"code"`
		Data model.ChapterImportResult `json:"data"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode chapter import response: %v", err)
	}
	if response.Code != recorder.Code {
		t.Errorf("response code reported premature success: %d", response.Code)
	}
	if !reflect.DeepEqual(response.Data.Imported, []string{"intro", "details"}) {
		t.Errorf("unexpected imported chapters: %v", response.Data.Imported)
	}
	if !strings.Contains(response.Data.Error, "malformed excerpt") {
		t.Errorf("decode error was swallowed: %q", response.Data.Error)
	}
	if response.Data.AuditStatus != "failed" {
		t.Errorf("failed import was audited as %q", response.Data.AuditStatus)
	}
	if response.Data.OpenReaders != 0 {
		t.Errorf("chapter readers remained open: %d", response.Data.OpenReaders)
	}
	if response.Data.PeakReaders != 1 {
		t.Errorf("reader lifetime crossed loop iterations, peak=%d", response.Data.PeakReaders)
	}
}
