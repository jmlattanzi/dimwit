package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmlattanzi/dimwit/api/home"
	"github.com/jmlattanzi/dimwit/api/task"
)

// Routes sets up the routes for our application
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Recoverer,
		middleware.Logger,
		middleware.SetHeader("Content-Type", "application/json"),
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api", home.HomeController{}.Routes())
		r.Mount("/api/task", task.TaskController{}.Routes())
	})

	return router
}

func main() {
	port := flag.String("port", ":8080", "Default network port")
	router := Routes()

	// might use this logger else
	infoLog := log.New(os.Stdout, "[ info ]\t", log.Ltime|log.Lshortfile)

	infoLog.Printf("Server starting on port %s", *port)
	log.Fatal(http.ListenAndServe(*port, router))
}
