package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/holedaemon/turnttable/internal/db/dbx"
	"github.com/holedaemon/turnttable/internal/web"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

type options struct {
	Addr    string            `env:"TURNTTABLE_ADDR"`
	Debug   bool              `env:"TURNTTABLE_DEBUG" envDefault:"false"`
	DSN     string            `env:"TURNTTABLE_DSN"`
	Admins  map[string]string `env:"TURNTTABLE_ADMINS"`
	Retries int               `env:"TURNTTABLE_DB_RETRIES" envDefault:"20"`
	Timeout time.Duration     `env:"TURNTTABLE_DB_TIMEOUT" envDefault:"10s"`
}

func main() {
	envOpts := env.Options{
		RequiredIfNoDef: true,
	}

	var opts options
	if err := env.ParseWithOptions(&opts, envOpts); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing env options: %s\n", err.Error())
		return
	}

	logger := ctxlog.New(opts.Debug)
	ctx := ctxlog.WithLogger(context.Background(), logger)
	connected := false
	var db *sql.DB

	for i := 0; i < opts.Retries && !connected; i++ {
		var err error

		db, err = sql.Open(dbx.Driver, opts.DSN)
		if err != nil {
			logger.Error("error connecting to database", zap.Int("attempt", i), zap.Error(err))
			time.Sleep(opts.Timeout)
			continue
		}

		if err = db.PingContext(ctx); err != nil {
			logger.Error("error pinging database", zap.Int("attempt", i), zap.Error(err))
			time.Sleep(opts.Timeout)
			continue
		}

		connected = true
	}

	if !connected {
		logger.Fatal("unable to connect to database after 20 attempts")
	}

	defer db.Close()

	srv := &web.Server{
		Addr:   opts.Addr,
		DB:     db,
		Admins: opts.Admins,
	}

	if err := srv.Run(ctx); err != nil {
		logger.Fatal("error running server", zap.Error(err))
	}
}
