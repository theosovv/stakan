// Command backfill loads historical funding/OI data from exchange REST APIs.
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
	exchange := flag.String("exchange", "", "exchange to backfill (binance|bybit|okx|bitget)")
	days := flag.Int("days", 30, "history depth in days")
	direct := flag.Bool("direct", false, "write directly to ClickHouse, bypassing Redpanda")
	flag.Parse()

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service", "backfill")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	httpx.MountHealth(mux, "backfill")

	log.Info("backfill stub started", "exchange", *exchange, "days", *days, "direct", *direct)
	if err := httpx.Run(ctx, log, config.Get("HTTP_ADDR", ":8106"), mux); err != nil {
		log.Error("server failed", "err", err)
		os.Exit(1)
	}
}
