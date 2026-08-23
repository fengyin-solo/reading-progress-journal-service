package handler

import (
	"net/http"
	"strconv"

	"readinglog/internal/model"
	"readinglog/pkg/httpx"
)

func (s *Server) registerProgressSyncRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/progress-sync/{bookID}", s.startProgressSync)
	mux.HandleFunc("POST /api/progress-sync/{id}/retry", s.retryProgressSync)
	mux.HandleFunc("POST /api/progress-sync/{id}/callback", s.completeProgressSync)
	mux.HandleFunc("GET /api/progress-sync/{id}", s.getProgressSync)
}

func (s *Server) startProgressSync(w http.ResponseWriter, r *http.Request) {
	httpx.Created(w, s.svc.StartProgressSync(r.PathValue("bookID")))
}

func (s *Server) retryProgressSync(w http.ResponseWriter, r *http.Request) {
	job, err := s.svc.RetryProgressSync(r.PathValue("id"))
	if err != nil { writeServiceError(w, err); return }
	httpx.OK(w, job)
}

func (s *Server) completeProgressSync(w http.ResponseWriter, r *http.Request) {
	version, err := strconv.Atoi(r.URL.Query().Get("version"))
	if err != nil { httpx.BadRequest(w, "version 不合法"); return }
	job, err := s.svc.CompleteProgressSync(r.PathValue("id"), version, model.SyncProcessing)
	if err != nil { writeServiceError(w, err); return }
	httpx.OK(w, job)
}

func (s *Server) getProgressSync(w http.ResponseWriter, r *http.Request) {
	detail, feed, notices, err := s.svc.ProgressSyncViews(r.PathValue("id"))
	if err != nil { writeServiceError(w, err); return }
	httpx.OK(w, map[string]interface{}{"detail": detail, "feed": feed, "notices": notices})
}
