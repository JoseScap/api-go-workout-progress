package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ExerciseRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get all"))
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create"))
	})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		w.Write([]byte("Get by id " + idParam))
	})
	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		w.Write([]byte("Put by id " + idParam))
	})
	r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		w.Write([]byte("Patch by id " + idParam))
	})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		w.Write([]byte("Delete by id " + idParam))
	})
	return r
}
