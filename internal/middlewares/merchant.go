package middlewares

import (
	"context"
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/service"
	mycontext "github.com/MamangRust/paymentgatewaygraphql/pkg/context"
)

func ApiKeyMiddleware(merchant service.MerchantService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-Api-Key")
			if apiKey == "" {
				http.Error(w, "API Key is required", http.StatusUnauthorized)
				return
			}

			_, err := merchant.FindByApiKey(apiKey)
			if err != nil {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), mycontext.ApiKeyContextKey, apiKey)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
