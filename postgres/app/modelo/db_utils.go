package modelo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

type ParametrosDB struct {
	DbDriver string
	DbSource string
}

var Parametros ParametrosDB

func LoadDbParameters() {

	// Open our jsonFile
	jsonFile, err := os.Open("./resources/properties_db.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var p ParametrosDB
	json.Unmarshal(byteValue, &p)
	defer jsonFile.Close()
	Parametros = p
	//fmt.Println(Parametros.DbDriver)

}
