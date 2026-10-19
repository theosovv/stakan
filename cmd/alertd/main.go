// Command alertd evaluates alert rules against the market-data stream.
package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/theosovv/stakan/pkg/config"
	"github.com/theosovv/stakan/pkg/httpx"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service", "alertd")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	httpx.MountHealth(mux, "alertd")

	log.Info("alertd stub started")
	if err := httpx.Run(ctx, log, config.Get("HTTP_ADDR", ":8103"), mux); err != nil {
		log.Error("server failed", "err", err)
		os.Exit(1)
	}
}
