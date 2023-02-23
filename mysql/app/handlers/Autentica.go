package handlers

import (
	"html/template"
	"net/http"

	"app/autentica"
)

func Login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/login.html")

}

func AutenticaAd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario := r.Form.Get("usuario")
	senha := r.Form.Get("senha")
	campos := map[string]interface{}{
		"usuario": usuario,
	}
	if autentica.Autentica(usuario, senha) {
		outputHTML(w, "./static/altera.html", campos)
	} else {
		http.ServeFile(w, r, "./static/login.html")
	}
}

func Autentica(usuario, senha string) {
	panic("unimplemented")
}

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
