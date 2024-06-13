package home

import (
	"map/internal/ai_integration"
	"net/http"

	"github.com/leapkit/core/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())

	err := rw.Render("home/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Reveal(w http.ResponseWriter, r *http.Request) {
	location := r.FormValue("location")
	aiGenerator := r.Context().Value("aiGenerator").(ai_integration.Service)

	response, err := aiGenerator.Generate(location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rw := render.FromCtx(r.Context())
	rw.Set("result", response)
	err = rw.RenderClean("home/result.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
