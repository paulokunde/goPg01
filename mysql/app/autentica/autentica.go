package autentica

import (
	"fmt"
	"log"

	auth "github.com/korylprince/go-ad-auth/v3"
)

func Autentica(usuario string, senha string) bool {
	adConfig := GetConfig()
	status, errr := auth.Authenticate(adConfig, usuario, senha)
	if errr != nil {
		log.Println("Erro ao Autenticar...")
		fmt.Printf("errr.Error(): %v\n", errr)
		return false
	}
	if !status {
		log.Println("Falha ao Autenticar....")
		return false
	}
	log.Println("Autenticado com Sucesso!")

	return true
}
