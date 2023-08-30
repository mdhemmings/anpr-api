package auth

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("Bearer token missing or invalid")
			return
		}

		token := strings.TrimPrefix(authorizationHeader, "Bearer ")

		// Validate the token (you can implement your own token validation logic here)
		if token != os.Getenv("APIKEY") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("Invalid bearer token")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for all requests
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with your desired origin(s)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continue processing the request
		next.ServeHTTP(w, r)
	})
}
