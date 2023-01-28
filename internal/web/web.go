package web

import (
	"context"
	"database/sql"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/holedaemon/turnttable/internal/web/mid"
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

func (s *Server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := chi.NewRouter()

	logger := ctxlog.FromContext(ctx)
	r.Use(mid.Logger(logger))
	r.Use(mid.Recoverer)

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
