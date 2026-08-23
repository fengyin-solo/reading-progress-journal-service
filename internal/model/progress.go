package model

import (
	"strings"
	"time"
)

// Progress 阅读进度记录。
type Progress struct {
	ID          string    `json:"id"`
	BookID      string    `json:"book_id"`
	CurrentPage int       `json:"current_page"`
	Percentage  float64   `json:"percentage"`
	RecordedAt  time.Time `json:"recorded_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// Validate 校验进度字段并规范化。
func (p *Progress) Validate() error {
	p.BookID = strings.TrimSpace(p.BookID)
	if p.BookID == "" {
		return NewValidationError("book_id", "书籍 ID 不能为空")
	}
	if p.CurrentPage < 0 {
		return NewValidationError("current_page", "当前页码不能为负数")
	}
	if p.Percentage < 0 || p.Percentage > 100 {
		return NewValidationError("percentage", "进度百分比必须在 0 到 100 之间")
	}
	if p.RecordedAt.IsZero() {
		p.RecordedAt = time.Now()
	}
	return nil
}

// ProgressFilter 进度查询筛选条件。
type ProgressFilter struct {
	BookID string
}

// Match 判断进度记录是否命中筛选条件。
func (f ProgressFilter) Match(p *Progress) bool {
	if f.BookID != "" && p.BookID != f.BookID {
		return false
	}
	return true
}
