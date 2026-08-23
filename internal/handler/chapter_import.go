package handler

import (
	"net/http"

	"readinglog/internal/service"
	"readinglog/pkg/httpx"
)

func ChapterImportScenario(w http.ResponseWriter, r *http.Request) {
	result, err := service.RunChapterImport()
	if err != nil {
		httpx.JSON(w, http.StatusBadRequest, 400, err.Error(), result)
		return
	}
	httpx.OK(w, result)
}
