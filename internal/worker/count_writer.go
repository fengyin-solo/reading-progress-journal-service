package worker

import "readinglog/internal/store"

type CountWriter struct {
	index *store.ReadIndex
}

func NewCountWriter(index *store.ReadIndex) *CountWriter {
	return &CountWriter{index: index}
}

func (w *CountWriter) ApplyMany(start <-chan struct{}, ready chan<- struct{}, done chan<- struct{}) {
	close(ready)
	<-start
	for i := 0; i < 100000; i++ {
		w.index.Increment("book-A")
	}
	close(done)
}
