package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gorilla/mux"
)

type Head struct {
	Title       string
	Description string
}

type Response struct {
	Head Head
	Main map[string]any
}

type Handlers struct {
	Renderer multitemplate.Renderer
}

func (h *Handlers) BikesHandler(w http.ResponseWriter, r *http.Request) {
	ren := h.Renderer.Instance("bike", Response{Main: make(map[string]any)})

	// if !bodyAllowedForStatus(code) {
	// 	r.WriteContentType(c.Writer)
	// 	c.Writer.WriteHeaderNow()
	// 	return
	// }

	w.WriteHeader(http.StatusOK)

	if err := ren.Render(w); err != nil {
		fmt.Println(err)
	}
}

func (h *Handlers) AgendaHandler(w http.ResponseWriter, r *http.Request) {
	ren := h.Renderer.Instance("agenda", Response{Main: make(map[string]any)})

	// if !bodyAllowedForStatus(code) {
	// 	r.WriteContentType(c.Writer)
	// 	c.Writer.WriteHeaderNow()
	// 	return
	// }

	w.WriteHeader(http.StatusOK)

	if err := ren.Render(w); err != nil {
		fmt.Println(err) // todo
	}
}

func (h *Handlers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	ren := h.Renderer.Instance("index", Response{Main: make(map[string]any)})

	// if !bodyAllowedForStatus(code) {
	// 	r.WriteContentType(c.Writer)
	// 	c.Writer.WriteHeaderNow()
	// 	return
	// }

	w.WriteHeader(http.StatusOK)

	if err := ren.Render(w); err != nil {
		fmt.Println(err)
	}
}

func createMyRender() multitemplate.Renderer {
	commonTemplates := []string{
		"./http/views/layouts/layout.html",
		"./http/views/partials/footer.html",
		"./http/views/partials/header.html",
		"./http/views/layouts/head.html",
		"./http/views/layouts/templates.html",
	}

	renderer := multitemplate.NewRenderer()
	renderer.AddFromFiles("index", append(commonTemplates, "./http/views/main/index.html")...)
	renderer.AddFromFiles("bike", append(commonTemplates, "./http/views/main/bike.html")...)
	renderer.AddFromFiles("agenda", append(commonTemplates, "./http/views/main/agenda.html")...)

	return renderer
}

func main() {
	h := &Handlers{Renderer: createMyRender()}

	r := mux.NewRouter()
	r.HandleFunc("/bikes/{id:[0-9]+}", h.BikesHandler)
	r.HandleFunc("/", h.IndexHandler)
	r.HandleFunc("/agenda", h.AgendaHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
