package handler

import (
	"net/http"

	"readinglog/internal/model"
	"readinglog/pkg/httpx"
)

func (s *Server) registerShelfRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/shelves", s.createShelf)
	mux.HandleFunc("GET /api/shelves", s.listShelves)
	mux.HandleFunc("GET /api/shelves/{id}", s.getShelf)
	mux.HandleFunc("PUT /api/shelves/{id}", s.updateShelf)
	mux.HandleFunc("DELETE /api/shelves/{id}", s.deleteShelf)
}

type shelfRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *Server) createShelf(w http.ResponseWriter, r *http.Request) {
	var req shelfRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	sh, err := s.svc.CreateShelf(model.Shelf{Name: req.Name, Description: req.Description})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, sh)
}

func (s *Server) listShelves(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.ShelfFilter{Keyword: r.URL.Query().Get("keyword")}
	items, total, err := s.svc.ListShelves(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getShelf(w http.ResponseWriter, r *http.Request) {
	sh, err := s.svc.GetShelf(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, sh)
}

func (s *Server) updateShelf(w http.ResponseWriter, r *http.Request) {
	var req shelfRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	sh, err := s.svc.UpdateShelf(r.PathValue("id"), model.Shelf{Name: req.Name, Description: req.Description})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, sh)
}

func (s *Server) deleteShelf(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteShelf(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
