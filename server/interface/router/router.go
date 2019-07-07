package router

import (
	"net/http"
	"todo/server/interface/registry"
)

func get(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.ServeHTTP(w, r)
		} else {
			http.Error(w, "invalid HTTP request.", http.StatusMethodNotAllowed)
		}
	}
}

func post(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.ServeHTTP(w, r)
		} else {
			http.Error(w, "invalid HTTP request.", http.StatusMethodNotAllowed)
		}
	}
}

func put(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preflightedHandler(w, r)

		if r.Method == http.MethodPut {
			handler.ServeHTTP(w, r)
		} else {
			http.Error(w, "invalid HTTP request.", http.StatusMethodNotAllowed)
		}
	}
}

func preflightedHandler(w http.ResponseWriter, r *http.Request) {
	addCorsHeader(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func addCorsHeader(w http.ResponseWriter) {
	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

func Init(r registry.HandlerRegistry) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/items", get(r.ItemHandler.Index))
	mux.HandleFunc("/items/create", post(r.ItemHandler.Create))
	mux.HandleFunc("/items/close", put(r.ItemHandler.Close))
	mux.HandleFunc("/items/open", put(r.ItemHandler.Open))

	return mux
}
