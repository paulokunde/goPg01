package testes

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"app/modelo"
)

func createRandomMarca(t *testing.T) modelo.Marca {
	var testQueries *modelo.Queries

	nome := "HP do Brasil"
	fmt.Println(nome)
	marca, err := testQueries.CreateMarca(context.Background(), nome)

	require.NoError(t, err)
	require.NotEmpty(t, marca)
	require.NotEmpty(t, marca.ID)
	require.Equal(t, nome, marca.Nome)
	return marca

}

func testCreateMarca(t *testing.T) {
	createRandomMarca(t)
}

func TestGetMarca(t *testing.T) {
	var testQueries *modelo.Queries
	marcaRandomCreated := createRandomMarca(t)
	marcaFinded, err := testQueries.GetMarcaById(context.Background(), marcaRandomCreated.ID)
	if err != nil {
		log.Fatal("Erro ao testar marca", err)
	}
	fmt.Printf(marcaFinded.Nome)

}
