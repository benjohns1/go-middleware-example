package middleware

import "net/http"

// Factory generates an http handler func in a middleware chain
type Factory func(next http.HandlerFunc) http.HandlerFunc

// Chain generates an http handler recursively from a list of middleware factory functions
func Chain(chain ...Factory) http.Handler {
	return http.HandlerFunc(recurseChain(chain))
}

func recurseChain(chain []Factory) http.HandlerFunc {
	if len(chain) <= 0 {
		return func(_ http.ResponseWriter, _ *http.Request) {}
	}
	return chain[0](recurseChain(chain[1:]))
}
