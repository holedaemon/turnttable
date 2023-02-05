package web

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
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
Admin insert page
Admin bulk insert page
Admin delete page
Admin edit page
*/

var mediumMap map[string]models.Medium = map[string]models.Medium{
	"cd":       models.MediumCD,
	"vinyl":    models.MediumVinyl,
	"cassette": models.MediumCassette,
}

//go:embed static
var static embed.FS

var staticDir fs.FS

func init() {
	var err error
	staticDir, err = fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := chi.NewRouter()

	logger := ctxlog.FromContext(ctx)
	r.Use(Logger(logger))
	r.Use(s.recoverer)

	r.Get("/", s.index)
	r.Get("/about", s.about)
	r.Route("/admin", s.routeAdmin)

	r.NotFound(s.notFound)

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.FS(staticDir))))
	r.Handle("/favicon.ico", http.RedirectHandler("/static/favicon.ico", http.StatusFound))

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

func (s *Server) about(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.Aboutpage{})
}
