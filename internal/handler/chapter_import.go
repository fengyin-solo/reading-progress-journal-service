package handler

import (
	"net/http"

	"readinglog/internal/service"
	"readinglog/pkg/httpx"
)

func ChapterImportScenario(w http.ResponseWriter, r *http.Request) {
	result, _ := service.RunChapterImport()
	httpx.OK(w, result)
}
