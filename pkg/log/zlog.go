package log

import (
	"github.com/glats/go-ms/pkg/api"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"context"
	"github.com/rs/zerolog"
	"time"
)

// Service represents zerolog logger
type Service struct {
	l zerolog.Logger
}


// FromContext fetches context from logger
func FromContext(ctx context.Context) *zerolog.Event {
	return ctx.Value(api.ContextKey("logger")).(*zerolog.Event)
}

// New instantiates new zero logger
func New() Service {
	z := zerolog.New(os.Stdout)
	zerolog.DurationFieldInteger = true
	return Service{l: z}
}

// Middleware represents logging middleware
func (s Service) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		rec := httptest.NewRecorder()

		req, _ := httputil.DumpRequest(r, r.Method != http.MethodGet)

		ctx := r.Context()

		logger := s.l.Info()
		logger.Timestamp().Str("path", r.URL.EscapedPath()).Str("ip", r.RemoteAddr).
			Interface("request_id", ctx.Value(middleware.RequestIDKey)).Bytes("request", req)

		var body string
		var queries []string

		defer func(begin time.Time) {
			status := ww.Status()
			logger.Int64("took", time.Since(begin).Milliseconds()).Int("status", status).
				Strs("queries", queries)

			if status != http.StatusNotFound {
				logger.Str("response", body)
			}

			if status >= 500 || ww.Status() < 600{
				logger.Send()
			}

		}(time.Now())
		ctx = context.WithValue(ctx, api.ContextKey("logger"), logger)
		ctx = context.WithValue(ctx, api.ContextKey("query"), &queries)

		next.ServeHTTP(rec, r.WithContext(ctx))

		// this copies the recorded response to the response writer
		for k, v := range rec.Header() {
			ww.Header()[k] = v
		}

		body = rec.Body.String()
		ww.WriteHeader(rec.Code)
		rec.Body.WriteTo(ww)
	})
}