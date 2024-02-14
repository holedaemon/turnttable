package web

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
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

var mediumMap map[string]models.Medium = map[string]models.Medium{
	"cd":       models.MediumCD,
	"vinyl":    models.MediumVinyl,
	"cassette": models.MediumCassette,
}

// Server is web server responsible for serving turnttable.
type Server struct {
	addr   string
	db     *sql.DB
	admins map[string]string
}

// New creates a new server from the given options.
func New(opts ...Option) (*Server, error) {
	s := &Server{}

	for _, o := range opts {
		o(s)
	}

	if s.addr == "" {
		return nil, fmt.Errorf("web: missing address")
	}

	if s.db == nil {
		return nil, fmt.Errorf("web: missing db")
	}

	if s.admins == nil {
		return nil, fmt.Errorf("web: missing admins")
	}

	return s, nil
}

// Run registers routes and runs a server.
func (s *Server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := chi.NewRouter()

	logger := ctxlog.FromContext(ctx)
	r.Use(injectLogger(logger))
	r.Use(s.recoverer)

	r.Get("/", s.index)
	r.Route("/admin", s.routeAdmin)

	r.NotFound(s.notFound)

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.FS(staticDir))))
	r.Handle("/favicon.ico", http.RedirectHandler("/static/favicon.ico", http.StatusFound))

	srv := http.Server{
		Addr:        s.addr,
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
		rows, err = models.Records().All(ctx, s.db)
	} else {
		rows, err = models.Records(qm.Where("MEDIUM = ?", filter)).All(ctx, s.db)
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
