package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/phsym/console-slog"
)

func newLogger(env string) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	if env != "production" {
		return slog.New(console.NewHandler(os.Stderr, &console.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	} else {
		return slog.New(slog.NewJSONHandler(os.Stderr, opts))
	}
}

func (app *application) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		arw := &AdvancedResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(arw, r)

		app.Logger.Info("Request",
			slog.String("Method", r.Method),
			slog.Int("Status", arw.statusCode),
			slog.String("Path", r.URL.Path),
			slog.Int("Bytes", arw.bytes),
			slog.Duration("time", time.Since(start)),
		)
	})
}
