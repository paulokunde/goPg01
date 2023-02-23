package autentica

import (
	"log"

	auth "github.com/korylprince/go-ad-auth/v3"
)

var adConfig *auth.Config

func ConectaAD() (*auth.Conn, error) {
	log.Print("-------Connectando no AD------------")
	conn, err := GetConfig().Connect()
	if err != nil {
		log.Printf("Erro ao conectar no AD")
	}
	return conn, err
}
