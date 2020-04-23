package task

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// TaskController is a wrapper
type TaskController struct {
}

// Routes returns all of the routes for this api
func (tc TaskController) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", GetTask)

	return router
}

// GetTask returns a task
func GetTask(w http.ResponseWriter, r *http.Request) {
	task := &Task{
		Title:   "task",
		Content: "Content of the task",
	}

	_ = json.NewEncoder(w).Encode(task)
}
