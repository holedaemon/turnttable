package web

import (
	"database/sql"
	"encoding/csv"
	"errors"
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

func (s *Server) routeAdmin(r chi.Router) {
	r.Use(s.adminAuth)

	r.Get("/insert", s.adminInsert)
	r.Post("/insert", s.postAdminInsert)

	r.Get("/insert/bulk", s.adminBulkInsert)
	r.Post("/insert/bulk", s.postAdminBulkInsert)

	r.Get("/edit/{id}", s.adminEdit)
	r.Post("/edit/{id}", s.postAdminEdit)

	r.Get("/delete/{id}", s.adminDelete)
	r.Post("/delete/{id}", s.postAdminDelete)
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

func (s *Server) adminInsert(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.AdminPage{})
}

func (s *Server) postAdminInsert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		ctxlog.Error(ctx, "error parsing form", zap.Error(err))
		s.internalError(w)
		return
	}

	title := r.FormValue("title")
	medium := mediumMap[strings.ToLower(r.FormValue("medium"))]

	exists, err := models.Records(
		qm.Where("title ILIKE ? AND medium = ?", title, medium),
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
		s.badRequest(w, "Release time formatting is bad")
		return
	}

	pt, err := time.Parse("2006-01-02", pur)
	if err != nil {
		ctxlog.Error(ctx, "error parsing purchase time", zap.Error(err))
		s.badRequest(w, "Purchased time formatting is bad")
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
		Medium:    medium,
	}

	if err := record.Insert(ctx, s.DB, boil.Infer()); err != nil {
		ctxlog.Error(ctx, "error inserting", zap.Error(err))
		s.internalError(w)
		return
	}

	s.statusPage(w, "New Record", "wow okay cool", "Record has been inserted")
}

func (s *Server) adminBulkInsert(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.AdminBulkInsertPage{})
}

func (s *Server) postAdminBulkInsert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("records")
	if err != nil {
		ctxlog.Error(ctx, "retrieving form file", zap.Error(err))
		s.internalError(w)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		ctxlog.Error(ctx, "reading every row", zap.Error(err))
		s.internalError(w)
		return
	}

	for idx, row := range rows {
		if idx == 0 {
			if !checkHeaderOrder(row) {
				s.badRequest(w, "Wrong CSV row order. Fix it and try again, idiot")
				return
			}

			continue
		}

		released := row[5]
		pur := row[6]

		rt, err := time.Parse("2006-01-02", released)
		if err != nil {
			ctxlog.Error(ctx, "error parsing release time", zap.Error(err))
			s.badRequest(w, "Release time formatting is bad")
			return
		}

		pt, err := time.Parse("2006-01-02", pur)
		if err != nil {
			ctxlog.Error(ctx, "error parsing purchase time", zap.Error(err))
			s.badRequest(w, "Purchased time formatting is bad")
			return
		}

		record := &models.Record{
			Title:     row[0],
			Artist:    row[1],
			Label:     row[2],
			CN:        row[3],
			Genre:     row[4],
			Released:  rt,
			Purchased: null.TimeFrom(pt),
			Medium:    mediumMap[strings.ToLower(row[7])],
		}

		if err := record.Insert(ctx, s.DB, boil.Infer()); err != nil {
			ctxlog.Error(ctx, "error inserting record into database", zap.Error(err))
			s.internalError(w)
			return
		}
	}

	s.statusPage(w, "yay dood", "Wow cool okay", "Bulk insert has completed")
}

func (s *Server) adminEdit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()

	row, err := models.Records(qm.Where("id = ?", id)).One(ctx, s.DB)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			s.badRequest(w, "No record by that \"id\" exists")
			return
		}

		ctxlog.Error(ctx, "error getting record", zap.Error(err), zap.String("id", id))
		s.internalError(w)
		return
	}

	templates.WritePageTemplate(w, &templates.AdminEditPage{
		Record: row,
	})
}

func (s *Server) postAdminEdit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		ctxlog.Error(ctx, "error parsing form", zap.Error(err))
		s.internalError(w)
		return
	}

	title := r.FormValue("title")
	medium := mediumMap[strings.ToLower(r.FormValue("medium"))]

	exists, err := models.Records(
		qm.Where("title ILIKE ? AND medium = ? AND id != ?", title, medium, id),
	).Exists(ctx, s.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying for record", zap.Error(err))
		s.internalError(w)
		return
	}

	if exists {
		s.statusPage(w, "lol", "Your title & medium conflict with another record!!", "Silly goose")
		return
	}

	released := r.FormValue("release")
	pur := r.FormValue("purchased")

	rt, err := time.Parse("2006-01-02", released)
	if err != nil {
		ctxlog.Error(ctx, "error parsing release time", zap.Error(err))
		s.badRequest(w, "Release time formatting is bad")
		return
	}

	pt, err := time.Parse("2006-01-02", pur)
	if err != nil {
		ctxlog.Error(ctx, "error parsing purchase time", zap.Error(err))
		s.badRequest(w, "Purchased time formatting is bad")
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
		Medium:    medium,
	}

	if err := record.Insert(ctx, s.DB, boil.Infer()); err != nil {
		ctxlog.Error(ctx, "error inserting", zap.Error(err))
		s.internalError(w)
		return
	}

	s.statusPage(w, "Updated Record", "wow okay cool", "Record has been updated")
}

func (s *Server) adminDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()

	row, err := models.Records(qm.Where("id = ?", id)).One(ctx, s.DB)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			s.badRequest(w, "No record by that \"id\" exists")
			return
		}

		ctxlog.Error(ctx, "error getting record", zap.Error(err), zap.String("id", id))
		s.internalError(w)
		return
	}

	templates.WritePageTemplate(w, &templates.AdminDeletePage{
		Record: row,
	})
}

func (s *Server) postAdminDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()

	row, err := models.Records(qm.Where("id = ?", id)).One(ctx, s.DB)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			s.badRequest(w, "No record by that \"id\" exists")
			return
		}

		ctxlog.Error(ctx, "error getting record", zap.Error(err), zap.String("id", id))
		s.internalError(w)
		return
	}

	if err := row.Delete(ctx, s.DB); err != nil {
		ctxlog.Error(ctx, "error deleting row", zap.Error(err))
		s.internalError(w)
		return
	}

	s.statusPage(w, ":3", "bye sis", "Record has been deleted")
}
