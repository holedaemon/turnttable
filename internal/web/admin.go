package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/holedaemon/turnttable/internal/web/templates"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func (s *Server) routeAdmin(r chi.Router) {
	r.Use(middleware.NoCache)
	r.Use(s.adminAuth)

	r.Get("/", s.admin)
	r.Post("/", s.postAdmin)
}

func (s *Server) adminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(s.Admins) == 0 {
			//TODO: Render an error page
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok {
			s.unauthorized(w, r)
			return
		}

		expected := s.Admins[user]
		if expected == "" || pass != expected {
			s.unauthorized(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) admin(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.AdminPage{})
}

func (s *Server) postAdmin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ctxlog.Error(r.Context(), "error parsing form", zap.Error(err))
		s.internalError(w, r)
		return
	}

	ctxlog.Debug(r.Context(), "", zap.Any("", r.PostForm))
}
