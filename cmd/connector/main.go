// Command connector streams one exchange's WebSocket feed into Redpanda
// as canonical events. One process per exchange: -exchange=binance.
package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/theosovv/stakan/pkg/config"
	"github.com/theosovv/stakan/pkg/httpx"
)

func main() {
	exchange := flag.String("exchange", "", "exchange to connect (binance|bybit|okx|bitget)")
	flag.Parse()

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service", "connector")

	if *exchange == "" {
		log.Error("-exchange flag is required")
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	httpx.MountHealth(mux, "connector")

	log.Info("connector stub started", "exchange", *exchange)
	if err := httpx.Run(ctx, log, config.Get("HTTP_ADDR", ":8101"), mux); err != nil {
		log.Error("server failed", "err", err)
		os.Exit(1)
	}
}
