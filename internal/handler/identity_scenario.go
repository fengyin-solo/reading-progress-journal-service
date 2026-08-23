package handler

import (
	"net/http"

	"readinglog/pkg/httpx"
)

func (s *Server) registerIdentityScenarioRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/identity-scenario", s.runIdentityScenario)
}

func (s *Server) runIdentityScenario(w http.ResponseWriter, _ *http.Request) {
	httpx.OK(w, s.svc.RunIdentityScenario())
}
