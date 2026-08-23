package model

import (
	"strings"
	"time"
)

// Shelf 书架分类。
type Shelf struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validate 校验书架字段并规范化。
func (s *Shelf) Validate() error {
	s.Name = strings.TrimSpace(s.Name)
	s.Description = strings.TrimSpace(s.Description)
	if s.Name == "" {
		return NewValidationError("name", "书架名称不能为空")
	}
	return nil
}

// ShelfFilter 书架查询筛选条件。
type ShelfFilter struct {
	Keyword string
}

// Match 判断书架是否命中筛选条件。
func (f ShelfFilter) Match(s *Shelf) bool {
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" &&
			!strings.Contains(strings.ToLower(s.Name), k) &&
			!strings.Contains(strings.ToLower(s.Description), k) {
			return false
		}
	}
	return true
}
