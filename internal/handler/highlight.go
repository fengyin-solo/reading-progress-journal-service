package handler

import (
	"net/http"

	"readinglog/internal/model"
	"readinglog/pkg/httpx"
)

func (s *Server) registerHighlightRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/highlights", s.createHighlight)
	mux.HandleFunc("GET /api/highlights", s.listHighlights)
	mux.HandleFunc("GET /api/highlights/{id}", s.getHighlight)
	mux.HandleFunc("PUT /api/highlights/{id}", s.updateHighlight)
	mux.HandleFunc("DELETE /api/highlights/{id}", s.deleteHighlight)
	mux.HandleFunc("POST /api/highlights/batch-delete", s.batchDeleteHighlights)
}

type highlightRequest struct {
	BookID  string `json:"book_id"`
	Content string `json:"content"`
	Page    int    `json:"page"`
	Chapter string `json:"chapter"`
}

func (s *Server) createHighlight(w http.ResponseWriter, r *http.Request) {
	var req highlightRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	h, err := s.svc.CreateHighlight(model.Highlight{
		BookID:  req.BookID,
		Content: req.Content,
		Page:    req.Page,
		Chapter: req.Chapter,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, h)
}

func (s *Server) listHighlights(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.HighlightFilter{
		BookID:  r.URL.Query().Get("book_id"),
		Chapter: r.URL.Query().Get("chapter"),
	}
	items, total, err := s.svc.ListHighlights(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getHighlight(w http.ResponseWriter, r *http.Request) {
	h, err := s.svc.GetHighlight(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, h)
}

func (s *Server) updateHighlight(w http.ResponseWriter, r *http.Request) {
	var req highlightRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	h, err := s.svc.UpdateHighlight(r.PathValue("id"), model.Highlight{
		BookID:  req.BookID,
		Content: req.Content,
		Page:    req.Page,
		Chapter: req.Chapter,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, h)
}

func (s *Server) deleteHighlight(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteHighlight(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type batchDeleteHighlightsRequest struct {
	IDs []string `json:"ids"`
}

func (s *Server) batchDeleteHighlights(w http.ResponseWriter, r *http.Request) {
	var req batchDeleteHighlightsRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	count, err := s.svc.BatchDeleteHighlights(req.IDs)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, map[string]int{"deleted": count})
}
