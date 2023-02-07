package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/holedaemon/turnttable/internal/web"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	rawDebug := os.Getenv("TURNTTABLE_DEBUG")
	debug := rawDebug != ""

	var (
		logger *zap.Logger
		err    error
	)

	if debug {
		logger, err = zap.NewDevelopment()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating development logger: %s\n", err.Error())
			return
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating production logger: %s\n", err.Error())
			return
		}
	}

	ctx := ctxlog.WithLogger(context.Background(), logger)

	addr := os.Getenv("TURNTTABLE_ADDR")
	if addr == "" {
		logger.Fatal("addr is blank")
		return
	}

	rawAdmins := os.Getenv("TURNTTABLE_ADMINS")
	if rawAdmins == "" {
		logger.Fatal("at least one admin account is required")
		return
	}

	adminPairs := strings.Split(rawAdmins, ",")
	admins := make(map[string]string)

	for _, a := range adminPairs {
		split := strings.Split(a, ":")

		if len(split) < 2 {
			logger.Fatal("admin missing password", zap.String("admin", split[0]))
			return
		}

		if len(split) >= 3 {
			logger.Fatal("admin has invalid password", zap.String("admin", split[0]))
			return
		}

		admins[split[0]] = split[1]
	}

	dsn := os.Getenv("TURNTTABLE_DSN")
	if dsn == "" {
		logger.Fatal("postgres DSN is blank")
		return
	}

	connected := false
	var db *sql.DB

	timeout := time.Second * 10

	for i := 0; i < 20 && !connected; i++ {
		var err error

		db, err = sql.Open("pgx", dsn)
		if err != nil {
			logger.Error("error connecting to database", zap.Int("attempt", i), zap.Error(err))
			time.Sleep(timeout)
			continue
		}

		if err = db.PingContext(ctx); err != nil {
			logger.Error("error pinging database", zap.Int("attempt", i), zap.Error(err))
			time.Sleep(timeout)
			continue
		}

		connected = true
	}

	if !connected {
		logger.Fatal("unable to connect to database after 20 attempts")
	}

	defer db.Close()

	srv := &web.Server{
		Addr:   addr,
		DB:     db,
		Admins: admins,
	}

	if err := srv.Run(ctx); err != nil {
		logger.Fatal("error running server", zap.Error(err))
	}
}
