package testes

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"app/modelo"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:pmscs.123@db:5432/infra?sslmode=disable"
)

var TestQueries *modelo.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect db", err)
	}
	fmt.Println("Conectado.....OK")
	TestQueries = modelo.New(conn)
	os.Exit(m.Run())
}
