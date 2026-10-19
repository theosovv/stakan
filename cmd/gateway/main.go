// Command gateway serves the public REST API and the WebSocket hub.
package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/theosovv/stakan/pkg/config"
	"github.com/theosovv/stakan/pkg/httpx"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service", "gateway")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	httpx.MountHealth(mux, "gateway")
	mux.HandleFunc("GET /api/v1/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "stakan api"})
	})

	log.Info("gateway stub started")
	if err := httpx.Run(ctx, log, config.Get("HTTP_ADDR", ":8080"), mux); err != nil {
		log.Error("server failed", "err", err)
		os.Exit(1)
	}
}
