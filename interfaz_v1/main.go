package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"time"
)

func main() {
	router := chi.NewRouter()
	server := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	router.Get("/api/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{ "message": "bar" }`))
	})

	FileServer(router)

	panic(server.ListenAndServe())
}

// FileServer is serving static files.
func FileServer(router *chi.Mux) {
	root := "./interfaz_seeker/dist"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}
