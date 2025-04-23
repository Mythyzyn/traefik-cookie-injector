package cookieinjector

import (
	"context"
	"net/http"
	"strings"
)

// Config holds the plugin configuration.
type Config struct{}

// CreateConfig creates the default plugin config.
func CreateConfig() *Config {
	return &Config{}
}

// CookieInjector is a Traefik middleware.
type CookieInjector struct {
	next http.Handler
	name string
}

// New creates a new CookieInjector middleware.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &CookieInjector{
		next: next,
		name: name,
	}, nil
}

// ServeHTTP modifies Set-Cookie headers.
func (c *CookieInjector) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	w := &cookieResponseWriter{ResponseWriter: rw}
	c.next.ServeHTTP(w, req)

	for _, cookie := range w.modifiedCookies {
		rw.Header().Add("Set-Cookie", cookie)
	}
}

type cookieResponseWriter struct {
	http.ResponseWriter
	modifiedCookies []string
}

func (w *cookieResponseWriter) WriteHeader(statusCode int) {
	original := w.Header()["Set-Cookie"]
	w.Header().Del("Set-Cookie")

	for _, cookie := range original {
		if !strings.Contains(strings.ToLower(cookie), "secure") {
			cookie += "; Secure"
		}
		if !strings.Contains(strings.ToLower(cookie), "httponly") {
			cookie += "; HttpOnly"
		}
		if !strings.Contains(strings.ToLower(cookie), "samesite") {
			cookie += "; SameSite=Strict"
		}
		w.modifiedCookies = append(w.modifiedCookies, cookie)
	}

	w.ResponseWriter.WriteHeader(statusCode)
}
