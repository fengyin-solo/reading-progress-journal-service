package service

import (
	"fmt"

	"readinglog/internal/audit"
	"readinglog/internal/model"
	"readinglog/internal/source"
)

// RunChapterImport 运行默认的章节导入场景（含坏章节用于演练失败路径）。
func RunChapterImport() (model.ChapterImportResult, error) {
	return runChapterImport([]model.ChapterInput{
		{ID: "intro", Body: "opening notes"},
		{ID: "details", Body: "chapter details"},
		{ID: "broken", Body: "MALFORMED"},
	})
}

// runChapterImport 执行单批章节导入：每轮读完立即关闭读取器，
// 任一章节解码或关闭失败即停止并保留错误，审计据实记 success/failed。
func runChapterImport(inputs []model.ChapterInput) (result model.ChapterImportResult, err error) {
	tracker := &source.ReaderTracker{}
	trail := &audit.ImportAudit{}
	defer func() {
		err = trail.Finish(err)
		result.AuditStatus = trail.Status()
		result.OpenReaders = tracker.OpenCount()
		result.PeakReaders = tracker.PeakCount()
		if err != nil {
			result.Error = err.Error()
		}
	}()

	for _, input := range inputs {
		reader := tracker.Open(input)
		_, readErr := reader.Read()
		closeErr := reader.Close()
		if readErr != nil {
			return result, fmt.Errorf("decode %s: %w", input.ID, readErr)
		}
		if closeErr != nil {
			return result, fmt.Errorf("close %s: %w", input.ID, closeErr)
		}
		result.Imported = append(result.Imported, input.ID)
	}
	return result, nil
}
