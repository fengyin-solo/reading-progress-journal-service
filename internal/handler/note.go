package handler

import (
	"net/http"

	"readinglog/internal/model"
	"readinglog/pkg/httpx"
)

func (s *Server) registerNoteRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/notes", s.createNote)
	mux.HandleFunc("GET /api/notes", s.listNotes)
	mux.HandleFunc("GET /api/notes/{id}", s.getNote)
	mux.HandleFunc("PUT /api/notes/{id}", s.updateNote)
	mux.HandleFunc("DELETE /api/notes/{id}", s.deleteNote)
}

type noteRequest struct {
	BookID  string `json:"book_id"`
	Content string `json:"content"`
	Page    int    `json:"page"`
}

func (s *Server) createNote(w http.ResponseWriter, r *http.Request) {
	var req noteRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	n, err := s.svc.CreateNote(model.Note{BookID: req.BookID, Content: req.Content, Page: req.Page})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, n)
}

func (s *Server) listNotes(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.NoteFilter{
		BookID:  r.URL.Query().Get("book_id"),
		Keyword: r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListNotes(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getNote(w http.ResponseWriter, r *http.Request) {
	n, err := s.svc.GetNote(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, n)
}

func (s *Server) updateNote(w http.ResponseWriter, r *http.Request) {
	var req noteRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	n, err := s.svc.UpdateNote(r.PathValue("id"), model.Note{BookID: req.BookID, Content: req.Content, Page: req.Page})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, n)
}

func (s *Server) deleteNote(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteNote(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
