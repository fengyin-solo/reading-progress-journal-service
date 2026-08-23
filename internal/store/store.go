// Package store 定义数据访问接口与内存实现。
package store

import (
	"errors"

	"readinglog/internal/model"
)

var (
	// ErrNotFound 表示记录不存在。
	ErrNotFound = errors.New("记录不存在")
	// ErrConflict 表示记录已存在或状态冲突。
	ErrConflict = errors.New("记录已存在或状态冲突")
)

// Store 聚合全部实体的数据访问方法，便于测试时替换实现。
type Store interface {
	// Book
	CreateBook(b *model.Book) error
	GetBook(id string) (*model.Book, error)
	ListBooks() []*model.Book
	UpdateBook(b *model.Book) error
	DeleteBook(id string) error

	// Shelf
	CreateShelf(s *model.Shelf) error
	GetShelf(id string) (*model.Shelf, error)
	ListShelves() []*model.Shelf
	UpdateShelf(s *model.Shelf) error
	DeleteShelf(id string) error

	// Note
	CreateNote(n *model.Note) error
	GetNote(id string) (*model.Note, error)
	ListNotes() []*model.Note
	UpdateNote(n *model.Note) error
	DeleteNote(id string) error

	// Highlight
	CreateHighlight(h *model.Highlight) error
	GetHighlight(id string) (*model.Highlight, error)
	ListHighlights() []*model.Highlight
	UpdateHighlight(h *model.Highlight) error
	DeleteHighlight(id string) error

	// Progress
	CreateProgress(p *model.Progress) error
	GetProgress(id string) (*model.Progress, error)
	ListProgresses() []*model.Progress
	UpdateProgress(p *model.Progress) error
	DeleteProgress(id string) error
}
