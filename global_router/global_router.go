package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleReq := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Farabi")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return // Handle preflight requests
		}
		mux.ServeHTTP(w, r)

	})
	return handleReq
}
