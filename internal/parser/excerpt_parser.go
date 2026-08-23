package parser

import "readinglog/internal/model"

type ExcerptParser struct {
	buffer []byte
}

func (p *ExcerptParser) Parse(id, text string) model.ExcerptBatch {
	if cap(p.buffer) < len(text) {
		p.buffer = make([]byte, len(text))
	}
	p.buffer = p.buffer[:len(text)]
	copy(p.buffer, text)
	return model.ExcerptBatch{ID: id, Payload: p.buffer}
}
