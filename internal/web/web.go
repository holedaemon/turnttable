package web

import (
	"context"
	"database/sql"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/holedaemon/turnttable/internal/db/models"
	"github.com/holedaemon/turnttable/internal/web/templates"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

// Server is web server responsible for serving turnttable.
type Server struct {
	Addr   string
	DB     *sql.DB
	Admins map[string]string
}

/*
-	Index page
-	Index page DB rows
-	Admin page
-	Admin page form POST
-	Admin page basic auth
-	Static
-	NotFound
*/

var mediumMap map[string]models.Medium = map[string]models.Medium{
	"cd":       models.MediumCD,
	"vinyl":    models.MediumVinyl,
	"cassette": models.MediumCassette,
}

func (s *Server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := chi.NewRouter()

	logger := ctxlog.FromContext(ctx)
	r.Use(Logger(logger))
	r.Use(s.recoverer)

	r.Get("/", s.index)

	r.Route("/admin", s.routeAdmin)

	srv := http.Server{
		Addr:        s.Addr,
		Handler:     r,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			ctxlog.Error(ctx, "error shutting down server", zap.Error(err))
		}
	}()

	ctxlog.Info(ctx, "web server listening", zap.String("addr", srv.Addr))
	return srv.ListenAndServe()
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	rf := r.URL.Query().Get("filter")
	filter := mediumMap[rf]
	if filter == "" {
		filter = "all"
	}

	ctx := r.Context()

	var (
		rows models.RecordSlice
		err  error
	)

	if filter == "all" {
		rows, err = models.Records().All(ctx, s.DB)
	} else {
		rows, err = models.Records(qm.Where("MEDIUM = ?", filter)).All(ctx, s.DB)
	}

	if err != nil {
		ctxlog.Error(ctx, "error fetching rows", zap.Error(err))
		s.internalError(w)
		return
	}

	templates.WritePageTemplate(w, &templates.IndexPage{
		Rows: rows,
	})
}

func (s *Server) unauthorized(w http.ResponseWriter, header bool) {
	if header {
		w.Header().Add("WWW-Authenticate", `Basic realm="turnttable"`)
	}

	w.WriteHeader(http.StatusUnauthorized)
	s.statusPage(w, "Oh Uh", "LOL HOW ARE YOU SO SMALL???", "You need special pants to open this door (Unauthorized)")
}

func (s *Server) internalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	templates.WritePageTemplate(w, &templates.StatusPage{
		PageTitle: "uh oh sisters",
		Header:    "idk man, I think it's your catalytic converter",
		Subtitle:  "An internal server error has occurred, try again later",
	})
}

func (s *Server) statusPage(w http.ResponseWriter, title, header, subtitle string) {
	templates.WritePageTemplate(w, &templates.StatusPage{
		PageTitle: title,
		Header:    header,
		Subtitle:  subtitle,
	})
}
