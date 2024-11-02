package routers

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"net/http"
	"workout/domain/models"
	"workout/infrastructure/database"

	"github.com/go-chi/chi/v5"
)

func ExerciseRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		exercises := models.Exercises{}
		if err := database.Connection.All(&exercises); err != nil {
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
		id, err := uuid.FromString(idParam)
		if err != nil {
			http.Error(w, "Invalid id format", http.StatusBadRequest)
			return
		}

		exercise := models.Exercise{}
		if err := database.Connection.Find(&exercise, id); err != nil {
			http.Error(w, "Exercise not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(exercise); err != nil {
			http.Error(w, "Failed to serialize exercise", http.StatusInternalServerError)
			return
		}
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
		id, err := uuid.FromString(idParam)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		exercise := models.Exercise{}
		if err := database.Connection.Find(&exercise, id); err != nil {
			http.Error(w, "Exercise not found", http.StatusNotFound)
			return
		}

		if err := database.Connection.Destroy(&exercise); err != nil {
			http.Error(w, "Failed to delete exercise", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	return r
}
