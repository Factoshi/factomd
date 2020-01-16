package controlpanel

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
)

type WebHandler interface {
	RegisterRoutes(router *mux.Router)
}

type webHandler struct {
}

func NewWebHandler() WebHandler {
	return &webHandler{}
}

func (handler *webHandler) RegisterRoutes(router *mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(handler.notFound)

	resourceDirectory, err := resourceDirectory()
	if err != nil {
		panic(err)
	}

	// handle static files
	cssHandler := http.FileServer(http.Dir(filepath.Join(resourceDirectory, "css")))
	jsHandler := http.FileServer(http.Dir(filepath.Join(resourceDirectory, "js")))
	imgHandler := http.FileServer(http.Dir(filepath.Join(resourceDirectory, "images")))
	fontsHandler := http.FileServer(http.Dir(filepath.Join(resourceDirectory, "fonts")))

	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", cssHandler))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", jsHandler))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imgHandler))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", fontsHandler))

	// register web endpoints
	router.HandleFunc("/", handler.indexHandler)
}

func (handler *webHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handle %s '%s' from %s", r.Method, r.URL.Path, r.RemoteAddr)

	resourceDirectory, err := resourceDirectory()
	if err != nil {
		fmt.Fprintf(w, handler.errorPage(err))
		return
	}

	t, err := template.ParseFiles(filepath.Join(resourceDirectory, "views/index.html"))
	if err != nil {
		fmt.Fprintf(w, handler.errorPage(err))
		return
	}

	t.Execute(w, nil)
}

func (handler *webHandler) notFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("page not found %s '%s' from %s", r.Method, r.URL.Path, r.RemoteAddr)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "custom 404")
}

func (handler *webHandler) errorPage(err error) string {
	return fmt.Sprintf("<h1>%s</h1><div>%v</div>", "error occured", err)
}

func resourceDirectory() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("no caller information to retrieve control panel resource directory")
		return "", fmt.Errorf("failed to load control panel resource")
	}
	fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	return filepath.Join(path.Dir(filename), "resources"), nil
}
