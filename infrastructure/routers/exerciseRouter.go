package routers

import (
	"encoding/json"
	"net/http"
	"workout/domain/models"

	"github.com/go-chi/chi/v5"
	"github.com/gobuffalo/pop/v6"
)

func ExerciseRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tx, err := pop.Connect("development")
		if err != nil {
			http.Error(w, "Cannot access DB", http.StatusInternalServerError)
			return
		}

		exercises := models.Exercises{}
		if err := tx.All(&exercises); err != nil {
			http.Error(w, "Cannot get exercises", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(exercises); err != nil {
			http.Error(w, "Failed to serialize exercises", http.StatusInternalServerError)
			return
		}
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
