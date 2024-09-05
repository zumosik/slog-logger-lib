package logger

import (
	"io"
	"log/slog"

	"github.com/zumosik/slog-logger-lib/slogpretty"
)

const (
	local = "local"
	dev   = "dev"
	prod  = "prod"
)

func SetupLogger(env string, writeTo ...io.Writer) *slog.Logger {
	var log *slog.Logger

	multi := io.MultiWriter(writeTo...)

	switch env {
	case local:
		log = setupPrettySlog(multi)
	case dev:
		log = slog.New(
			slog.NewTextHandler(multi, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case prod:
		log = slog.New(slog.NewTextHandler(multi, &slog.HandlerOptions{Level: slog.LevelError}))
	}

	return log
}

func setupPrettySlog(w io.Writer) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(w)

	return slog.New(handler)
}
