package service

import (
	"fmt"

	"readinglog/internal/audit"
	"readinglog/internal/model"
	"readinglog/internal/source"
)

func RunChapterImport() (result model.ChapterImportResult, err error) {
	tracker := &source.ReaderTracker{}
	trail := &audit.ImportAudit{}
	defer func() {
		err = trail.Finish(err)
		result.AuditStatus = trail.Status()
		result.OpenReaders = tracker.OpenCount()
		result.PeakReaders = tracker.PeakCount()
	}()

	inputs := []model.ChapterInput{
		{ID: "intro", Body: "opening notes"},
		{ID: "details", Body: "chapter details"},
		{ID: "broken", Body: "MALFORMED"},
	}
	for _, input := range inputs {
		reader := tracker.Open(input)
		defer reader.Close()
		if _, readErr := reader.Read(); readErr != nil {
			return result, fmt.Errorf("decode %s: %w", input.ID, readErr)
		}
		result.Imported = append(result.Imported, input.ID)
	}
	return result, nil
}
