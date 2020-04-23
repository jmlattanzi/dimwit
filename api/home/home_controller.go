package home

import (
	"net/http"

	"github.com/go-chi/chi"
)

type HomeController struct {
}

func (hc HomeController) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", home)

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
