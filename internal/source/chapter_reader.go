package source

import (
	"errors"

	"readinglog/internal/model"
)

type ReaderTracker struct {
	open int
	peak int
}

type ChapterReader struct {
	input   model.ChapterInput
	tracker *ReaderTracker
	closed  bool
}

func (t *ReaderTracker) Open(input model.ChapterInput) *ChapterReader {
	t.open++
	if t.open > t.peak {
		t.peak = t.open
	}
	return &ChapterReader{input: input, tracker: t}
}

func (r *ChapterReader) Read() (string, error) {
	if r.input.Body == "MALFORMED" {
		return "", errors.New("malformed excerpt")
	}
	return r.input.Body, nil
}

func (r *ChapterReader) Close() error {
	if r.closed {
		return nil
	}
	r.closed = true
	return nil
}

func (t *ReaderTracker) OpenCount() int { return t.open }
func (t *ReaderTracker) PeakCount() int { return t.peak }
