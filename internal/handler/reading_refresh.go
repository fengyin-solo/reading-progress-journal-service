package handler

import (
	"context"
	"net/http"
	"time"

	"readinglog/pkg/httpx"
)

func (s *Server) registerReadingRefreshRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/reading-refresh/{bookID}", s.startReadingRefresh)
	mux.HandleFunc("GET /api/reading-refresh/{id}", s.getReadingRefresh)
	mux.HandleFunc("POST /api/reading-refresh/shutdown", s.shutdownReadingRefresh)
}

func (s *Server) startReadingRefresh(w http.ResponseWriter, r *http.Request) {
	refresh := s.svc.StartReadingRefresh(context.Background(), r.PathValue("bookID"))
	httpx.Created(w, refresh)
}

func (s *Server) getReadingRefresh(w http.ResponseWriter, r *http.Request) {
	refresh, err := s.svc.GetReadingRefresh(r.PathValue("id"))
	if err != nil { writeServiceError(w, err); return }
	httpx.OK(w, refresh)
}

func (s *Server) shutdownReadingRefresh(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Millisecond)
	defer cancel()
	if err := s.svc.ShutdownReadingRefresh(ctx); err != nil { httpx.InternalError(w, err.Error()); return }
	httpx.OK(w, map[string]string{"status": "stopped"})
}
