package server

import (
	"github.com/go-chi/render"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//https://jonnylangefeld.com/blog/how-to-write-a-go-api-the-ultimate-guide#14-render-responses-with-go-chirender
//https://itnext.io/structuring-a-production-grade-rest-api-in-golang-c0229b3feedc
//https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/
// New methods
func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", healthCheck)
	return r
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	render.DefaultResponder(w, r, map[string]string{
		"message": "ok",
	})
	return
}
