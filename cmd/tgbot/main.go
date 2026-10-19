// Command tgbot delivers fired alerts to Telegram and manages bot-side rules.
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
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service", "tgbot")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	httpx.MountHealth(mux, "tgbot")

	log.Info("tgbot stub started")
	if err := httpx.Run(ctx, log, config.Get("HTTP_ADDR", ":8105"), mux); err != nil {
		log.Error("server failed", "err", err)
		os.Exit(1)
	}
}
