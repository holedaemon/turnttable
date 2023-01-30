package main

import (
	"context"
	"os"
	"strings"

	"github.com/holedaemon/turnttable/internal/web"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func main() {
	addr := os.Getenv("TURNTTABLE_ADDR")
	if addr == "" {
		panic("turnttable: addr not specified")
	}

	rawAdmins := os.Getenv("TURNTTABLE_ADMINS")
	adminPairs := strings.Split(rawAdmins, ",")
	admins := make(map[string]string)

	for _, a := range adminPairs {
		split := strings.Split(a, ":")

		if len(split) < 2 {
			panic("turntable: admin " + split[0] + " missing password")
		}

		if len(split) >= 3 {
			panic("turnttable: admin " + split[0] + " has invalid password")
		}

		admins[split[0]] = split[1]
	}

	srv := &web.Server{
		Addr:   addr,
		Admins: admins,
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("turnttable: error creating zap logger")
	}

	ctx := ctxlog.WithLogger(context.Background(), logger)

	if err := srv.Run(ctx); err != nil {
		panic("turnttable: error running server: " + err.Error())
	}
}
