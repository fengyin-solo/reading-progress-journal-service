package handler

import (
	"net/http"

	"readinglog/pkg/httpx"
)

func (s *Server) registerReadingDigestRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/reading-digests/{id}", s.buildReadingDigest)
	mux.HandleFunc("GET /api/reading-digests/{id}", s.getReadingDigest)
	mux.HandleFunc("GET /api/reading-digests", s.listReadingDigests)
}

func (s *Server) buildReadingDigest(w http.ResponseWriter, r *http.Request) {
	digest, err := s.svc.BuildReadingDigest(r.PathValue("id"), r.URL.Query().Get("title"))
	if err != nil { writeServiceError(w, err); return }
	httpx.Created(w, digest)
}

func (s *Server) getReadingDigest(w http.ResponseWriter, r *http.Request) {
	digest, err := s.svc.GetReadingDigest(r.PathValue("id"))
	if err != nil { writeServiceError(w, err); return }
	httpx.OK(w, digest)
}

func (s *Server) listReadingDigests(w http.ResponseWriter, _ *http.Request) {
	httpx.OK(w, s.svc.ListReadingDigests())
}
