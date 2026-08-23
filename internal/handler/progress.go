package handler

import (
	"net/http"

	"readinglog/internal/model"
	"readinglog/pkg/httpx"
)

func (s *Server) registerProgressRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/progresses", s.createProgress)
	mux.HandleFunc("GET /api/progresses", s.listProgresses)
	mux.HandleFunc("GET /api/progresses/{id}", s.getProgress)
	mux.HandleFunc("PUT /api/progresses/{id}", s.updateProgress)
	mux.HandleFunc("DELETE /api/progresses/{id}", s.deleteProgress)
}

type progressRequest struct {
	BookID      string  `json:"book_id"`
	CurrentPage int     `json:"current_page"`
	Percentage  float64 `json:"percentage"`
}

func (s *Server) createProgress(w http.ResponseWriter, r *http.Request) {
	var req progressRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	p, err := s.svc.CreateProgress(model.Progress{
		BookID:      req.BookID,
		CurrentPage: req.CurrentPage,
		Percentage:  req.Percentage,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, p)
}

func (s *Server) listProgresses(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.ProgressFilter{BookID: r.URL.Query().Get("book_id")}
	items, total, err := s.svc.ListProgresses(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getProgress(w http.ResponseWriter, r *http.Request) {
	p, err := s.svc.GetProgress(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, p)
}

func (s *Server) updateProgress(w http.ResponseWriter, r *http.Request) {
	var req progressRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	p, err := s.svc.UpdateProgress(r.PathValue("id"), model.Progress{
		BookID:      req.BookID,
		CurrentPage: req.CurrentPage,
		Percentage:  req.Percentage,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, p)
}

func (s *Server) deleteProgress(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteProgress(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
