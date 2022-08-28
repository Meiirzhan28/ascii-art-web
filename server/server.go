package server

import (
	"net/http"
	"text/template"

	"web/server/ascii"
)

type Server struct {
	mux *http.ServeMux
}

type Data struct {
	Name string
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Handle() *http.ServeMux {
	s.mux.HandleFunc("/", home)
	s.mux.HandleFunc("/ascii-art", asciiPage)
	return s.mux
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	temp, err := template.ParseFiles("template/html.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func asciiPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Header().Set("Allow", http.MethodPost)
		w.Write([]byte("Method not allowed!"))
		return
	}
	temp, err := template.ParseFiles("template/html.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	text := r.FormValue("input")
	style := r.FormValue("textstyle")
	str, err1 := ascii.Toascii(text, style)
	if err1 == 400 {
		w.WriteHeader(400)
		w.Header().Set("Allow", http.MethodPost)
		w.Write([]byte("Bad Request"))
		return
	}
	if err1 == 500 {
		w.WriteHeader(500)

		w.Write([]byte("Internal Server Error"))
		return
	}
	result := Data{
		Name: str,
	}
	temp.Execute(w, result)
}
