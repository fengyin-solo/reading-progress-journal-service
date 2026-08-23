package model

import (
	"strings"
	"time"
)

// Note 读书笔记。
type Note struct {
	ID        string    `json:"id"`
	BookID    string    `json:"book_id"`
	Content   string    `json:"content"`
	Page      int       `json:"page"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate 校验笔记字段并规范化。
func (n *Note) Validate() error {
	n.BookID = strings.TrimSpace(n.BookID)
	n.Content = strings.TrimSpace(n.Content)
	if n.BookID == "" {
		return NewValidationError("book_id", "书籍 ID 不能为空")
	}
	if n.Content == "" {
		return NewValidationError("content", "笔记内容不能为空")
	}
	if n.Page < 0 {
		return NewValidationError("page", "页码不能为负数")
	}
	return nil
}

// NoteFilter 笔记查询筛选条件。
type NoteFilter struct {
	BookID  string
	Keyword string
}

// Match 判断笔记是否命中筛选条件。
func (f NoteFilter) Match(n *Note) bool {
	if f.BookID != "" && n.BookID != f.BookID {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(n.Content), k) {
			return false
		}
	}
	return true
}
