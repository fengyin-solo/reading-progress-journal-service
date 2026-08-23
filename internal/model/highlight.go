package model

import (
	"strings"
	"time"
)

// Highlight 书摘。
type Highlight struct {
	ID        string    `json:"id"`
	BookID    string    `json:"book_id"`
	Content   string    `json:"content"`
	Page      int       `json:"page"`
	Chapter   string    `json:"chapter"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate 校验书摘字段并规范化。
func (h *Highlight) Validate() error {
	h.BookID = strings.TrimSpace(h.BookID)
	h.Content = strings.TrimSpace(h.Content)
	h.Chapter = strings.TrimSpace(h.Chapter)
	if h.BookID == "" {
		return NewValidationError("book_id", "书籍 ID 不能为空")
	}
	if h.Content == "" {
		return NewValidationError("content", "书摘内容不能为空")
	}
	if h.Page < 0 {
		return NewValidationError("page", "页码不能为负数")
	}
	return nil
}

// HighlightFilter 书摘查询筛选条件。
type HighlightFilter struct {
	BookID  string
	Chapter string
}

// Match 判断书摘是否命中筛选条件。
func (f HighlightFilter) Match(h *Highlight) bool {
	if f.BookID != "" && h.BookID != f.BookID {
		return false
	}
	if f.Chapter != "" && h.Chapter != f.Chapter {
		return false
	}
	return true
}
