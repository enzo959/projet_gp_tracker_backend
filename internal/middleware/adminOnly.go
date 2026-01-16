package middleware

import "net/http"

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value("role").(string)
		if !ok || role != "admin" {
			http.Error(w, "forbidden: admin only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
