package handler

import (
	"net/http"

	"readinglog/internal/model"
	"readinglog/pkg/httpx"
)

func (s *Server) registerBookRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/books", s.createBook)
	mux.HandleFunc("GET /api/books", s.listBooks)
	mux.HandleFunc("GET /api/books/{id}", s.getBook)
	mux.HandleFunc("PUT /api/books/{id}", s.updateBook)
	mux.HandleFunc("DELETE /api/books/{id}", s.deleteBook)
	mux.HandleFunc("POST /api/books/{id}/start", s.startReading)
	mux.HandleFunc("POST /api/books/{id}/finish", s.finishReading)
	mux.HandleFunc("POST /api/books/{id}/drop", s.dropReading)
}

type bookRequest struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	ISBN       string `json:"isbn"`
	Category   string `json:"category"`
	TotalPages int    `json:"total_pages"`
	Status     string `json:"status"`
}

func (s *Server) createBook(w http.ResponseWriter, r *http.Request) {
	var req bookRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	b, err := s.svc.CreateBook(model.Book{
		Title:      req.Title,
		Author:     req.Author,
		ISBN:       req.ISBN,
		Category:   req.Category,
		TotalPages: req.TotalPages,
		Status:     req.Status,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, b)
}

func (s *Server) listBooks(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.BookFilter{
		Category: r.URL.Query().Get("category"),
		Status:   r.URL.Query().Get("status"),
		Keyword:  r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListBooks(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getBook(w http.ResponseWriter, r *http.Request) {
	b, err := s.svc.GetBook(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, b)
}

func (s *Server) updateBook(w http.ResponseWriter, r *http.Request) {
	var req bookRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	b, err := s.svc.UpdateBook(r.PathValue("id"), model.Book{
		Title:      req.Title,
		Author:     req.Author,
		ISBN:       req.ISBN,
		Category:   req.Category,
		TotalPages: req.TotalPages,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, b)
}

func (s *Server) deleteBook(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteBook(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

func (s *Server) startReading(w http.ResponseWriter, r *http.Request) {
	b, err := s.svc.StartReading(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, b)
}

func (s *Server) finishReading(w http.ResponseWriter, r *http.Request) {
	b, err := s.svc.FinishReading(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, b)
}

func (s *Server) dropReading(w http.ResponseWriter, r *http.Request) {
	b, err := s.svc.DropReading(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, b)
}
