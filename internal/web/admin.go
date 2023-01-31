package web

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/holedaemon/turnttable/internal/db/models"
	"github.com/holedaemon/turnttable/internal/web/templates"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

var mediumMap map[string]models.Medium = map[string]models.Medium{
	"cd":       models.MediumCD,
	"vinyl":    models.MediumVinyl,
	"cassette": models.MediumCassette,
}

func (s *Server) routeAdmin(r chi.Router) {
	r.Use(s.adminAuth)

	r.Get("/", s.admin)
	r.Post("/", s.postAdmin)
}

func (s *Server) adminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(s.Admins) == 0 {
			s.unauthorized(w, false)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok {
			s.unauthorized(w, true)
			return
		}

		expected := s.Admins[user]
		if expected == "" || pass != expected {
			s.unauthorized(w, true)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) admin(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.AdminPage{})
}

func (s *Server) postAdmin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		ctxlog.Error(ctx, "error parsing form", zap.Error(err))
		s.internalError(w)
		return
	}

	title := r.FormValue("title")

	exists, err := models.Records(
		qm.Where("title ILIKE ?", title),
	).Exists(ctx, s.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying for record", zap.Error(err))
		s.internalError(w)
		return
	}

	if exists {
		s.statusPage(w, "lol", "Record already exists!!", "Silly goose")
		return
	}

	released := r.FormValue("release")
	pur := r.FormValue("purchased")

	rt, err := time.Parse("2006-01-02", released)
	if err != nil {
		ctxlog.Error(ctx, "error parsing release time", zap.Error(err))
		s.internalError(w)
		return
	}

	pt, err := time.Parse("2006-01-02", pur)
	if err != nil {
		ctxlog.Error(ctx, "error parsing purchase time", zap.Error(err))
		s.internalError(w)
		return
	}

	record := models.Record{
		Title:     r.FormValue("title"),
		Artist:    r.FormValue("artist"),
		Label:     r.FormValue("label"),
		CN:        r.FormValue("cn"),
		Genre:     r.FormValue("genre"),
		Released:  rt,
		Purchased: null.TimeFrom(pt),
		Medium:    mediumMap[strings.ToLower(r.FormValue("medium"))],
	}

	if err := record.Insert(ctx, s.DB, boil.Infer()); err != nil {
		ctxlog.Error(ctx, "error inserting", zap.Error(err))
		s.internalError(w)
		return
	}

	s.statusPage(w, "New Record", "wow okay cool", "Record has been inserted")
}
