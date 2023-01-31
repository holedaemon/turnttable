package web

// Thanks, Zik

import (
	"net/http"

	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

// Recoverer recovers from any panics that occurr during a request.
func (s *Server) recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				ctx := ctxlog.WithOptions(r.Context(), zap.AddStacktrace(zap.ErrorLevel))
				ctxlog.Error(ctx, "PANIC", zap.Any("value", rvr))

				s.internalError(w)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// Logger injects a logger into the Handler chain.
func Logger(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := ctxlog.WithLogger(r.Context(), logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
