package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/raire/raire/internal/api"
	"github.com/raire/raire/internal/app"
	"github.com/raire/raire/internal/platform/config"
	"github.com/raire/raire/internal/platform/logging"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "config:", err)
		os.Exit(1)
	}
	logger := logging.New(cfg.Env)

	if len(os.Args) > 1 && os.Args[1] == "seed-admin" {
		seedAdmin(cfg, logger)
		return
	}

	a, err := app.New(cfg)
	if err != nil {
		logger.Error("startup_failed", "error", err.Error())
		os.Exit(1)
	}
	defer func() { _ = a.Close() }()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go purgeLoop(ctx, a, logger)

	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           api.NewRouter(a, logger),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	go func() {
		logger.Info("server_started", "addr", cfg.Addr, "env", cfg.Env)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server_failed", "error", err.Error())
			stop()
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
	logger.Info("server_stopped")
}

func purgeLoop(ctx context.Context, a *app.App, logger *slog.Logger) {
	purge := func() {
		runCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		n, err := a.Directory.PurgeExpiredLocations(runCtx, time.Now().UTC())
		if err != nil {
			logger.Error("location_purge_failed", "error", err.Error())
			return
		}
		if n > 0 {
			logger.Info("location_purge_completed", "profiles", n)
		}
	}
	purge()
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			purge()
		}
	}
}

func seedAdmin(cfg config.Config, logger *slog.Logger) {
	if !cfg.DevMode {
		fmt.Fprintln(os.Stderr, "seed-admin is only available in dev/test")
		os.Exit(1)
	}
	email := os.Getenv("RAIRE_ADMIN_EMAIL")
	password := os.Getenv("RAIRE_ADMIN_PASSWORD")
	if email == "" || password == "" {
		fmt.Fprintln(os.Stderr, "set RAIRE_ADMIN_EMAIL and RAIRE_ADMIN_PASSWORD")
		os.Exit(1)
	}
	a, err := app.New(cfg)
	if err != nil {
		logger.Error("startup_failed", "error", err.Error())
		os.Exit(1)
	}
	defer func() { _ = a.Close() }()
	_, secret, codes, err := a.CreateOperator(context.Background(), email, password)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create operator:", err)
		os.Exit(1)
	}
	fmt.Println("operator created")
	fmt.Println("TOTP secret:", secret)
	fmt.Println("recovery codes (store once, then delete this output):")
	for _, code := range codes {
		fmt.Println(" ", code)
	}
}
