package model

import (
	"strings"
	"time"
)

// 书籍阅读状态常量。
const (
	BookWishlist = "wishlist" // 想读
	BookReading  = "reading"  // 在读
	BookFinished = "finished" // 读完
	BookDropped  = "dropped"  // 弃读
)

// Book 书籍信息。
type Book struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	Author     string     `json:"author"`
	ISBN       string     `json:"isbn"`
	Category   string     `json:"category"`
	TotalPages int        `json:"total_pages"`
	Status     string     `json:"status"`
	StartedAt  *time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// Validate 校验书籍字段并规范化。
func (b *Book) Validate() error {
	b.Title = strings.TrimSpace(b.Title)
	b.Author = strings.TrimSpace(b.Author)
	b.ISBN = strings.TrimSpace(b.ISBN)
	b.Category = strings.TrimSpace(b.Category)
	if b.Title == "" {
		return NewValidationError("title", "书名不能为空")
	}
	if b.Author == "" {
		return NewValidationError("author", "作者不能为空")
	}
	if b.TotalPages <= 0 {
		return NewValidationError("total_pages", "总页数必须大于 0")
	}
	if b.Status == "" {
		b.Status = BookWishlist
	}
	if b.Status != BookWishlist && b.Status != BookReading &&
		b.Status != BookFinished && b.Status != BookDropped {
		return NewValidationError("status", "书籍状态不合法")
	}
	return nil
}

// bookTransitions 书籍状态机。
var bookTransitions = map[string]map[string]bool{
	BookWishlist: {BookReading: true, BookDropped: true},
	BookReading:  {BookFinished: true, BookDropped: true},
}

// CanTransitionBookStatus 判断书籍状态是否允许从 from 流转到 to。
func CanTransitionBookStatus(from, to string) bool {
	if m, ok := bookTransitions[from]; ok {
		return m[to]
	}
	return false
}

// BookFilter 书籍查询筛选条件。
type BookFilter struct {
	Category string
	Status   string
	Keyword  string
}

// Match 判断书籍是否命中筛选条件。
func (f BookFilter) Match(b *Book) bool {
	if f.Category != "" && b.Category != f.Category {
		return false
	}
	if f.Status != "" && b.Status != f.Status {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" &&
			!strings.Contains(strings.ToLower(b.Title), k) &&
			!strings.Contains(strings.ToLower(b.Author), k) &&
			!strings.Contains(strings.ToLower(b.ISBN), k) {
			return false
		}
	}
	return true
}
