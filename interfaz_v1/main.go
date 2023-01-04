package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
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

	router.Post("/api/searchQuery", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`OK`))
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
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
