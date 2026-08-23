package handler

import (
	"net/http"
	"strconv"

	"readinglog/pkg/httpx"
)

func (s *Server) registerReadingPolicyRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/reading-policy/check", s.checkReadingPolicy)
	mux.HandleFunc("PUT /api/reading-policy/rules/{name}", s.setReadingRule)
}

func policyPages(r *http.Request) (int, error) {
	return strconv.Atoi(r.URL.Query().Get("pages"))
}

func (s *Server) checkReadingPolicy(w http.ResponseWriter, r *http.Request) {
	pages, err := policyPages(r)
	if err != nil { httpx.BadRequest(w, "pages 不合法"); return }
	if err := s.svc.CheckReadingPolicy(pages); err != nil { writeServiceError(w, err); return }
	httpx.OK(w, map[string]bool{"allowed": true})
}

func (s *Server) setReadingRule(w http.ResponseWriter, r *http.Request) {
	pages, err := policyPages(r)
	if err != nil { httpx.BadRequest(w, "pages 不合法"); return }
	s.svc.SetReadingRule(r.PathValue("name"), pages)
	httpx.OK(w, map[string]int{"pages": s.svc.ReadingRule(r.PathValue("name"))})
}
