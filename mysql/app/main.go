package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"app/handlers"
)

func main() {

	host := ":5443"
	r := chi.NewRouter()
	r.Get("/", handlers.Login)
	r.Post("/", handlers.AutenticaAd)

	fileServer := http.FileServer(http.Dir("./static/css/"))
	r.Handle("/css/*", http.StripPrefix("/css/", fileServer))

	err := http.ListenAndServeTLS(host, "certificado.crt", "certificado.key", r)
	if err != nil {
		log.Printf("Erro ao inicir serviço na porta:%s", host)
		fmt.Printf("err: %v\n", err)
	}

}
