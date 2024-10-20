package middleware

import (
	"fmt"
	httpInternal "github.com/rsgcata/gocommon/infrastructure/http"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const AccessLogMessage = "HTTP Request"

type HttpAccessLogger struct {
	nextHandler http.Handler
	logger      *slog.Logger
}

func NewHttpAccessLogger(next http.Handler, logger *slog.Logger) *HttpAccessLogger {
	return &HttpAccessLogger{next, logger}
}

func (accessLogger *HttpAccessLogger) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	logResponseWriter := httpInternal.NewResponseWriter(rw)
	timeBeforeServe := time.Now().UnixMilli()
	accessLogger.nextHandler.ServeHTTP(logResponseWriter, rq)
	timeAfterServe := time.Now().UnixMilli()

	accessLogger.logger.LogAttrs(
		rq.Context(),
		slog.LevelInfo,
		AccessLogMessage,
		slog.String("Client IP", strings.Split(rq.RemoteAddr, ":")[0]),
		slog.String("Method", rq.Method),
		slog.String("Host", rq.Host),
		slog.String("Path", rq.URL.Path),
		slog.String("Query", rq.URL.RawQuery),
		slog.String("Protocol", rq.Proto),
		slog.String("User Agent", rq.UserAgent()),
		slog.String("Response Status Code", strconv.Itoa(logResponseWriter.StatusCode())),
		slog.String(
			"Duration (s)",
			fmt.Sprintf("%.2f", float64(timeAfterServe-timeBeforeServe)/1000),
		),
	)
}
