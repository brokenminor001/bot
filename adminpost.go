package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var PostRasp string
var Datum string
var Typ string

const (
	host     = "95.174.94.157"
	port     = 5432
	user     = "asmolin"
	password = "oadnpvia"
	dbname   = "bot"
)

type DataInputPost struct {
	PostRaspisanie       string
	PostRaspisanieData   string
	PostRaspisanieVremya string
	PostRaspisanieName   string
}

func main() {

	http.HandleFunc("/postraspisanie", PostHandler)

	http.ListenAndServe(":8081", nil)

}
func PostHandler(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t DataInputPost
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	Typ := t.PostRaspisanie
	Datum := t.PostRaspisanieData
	Vremya := t.PostRaspisanieVremya
	Name := t.PostRaspisanieName
	fmt.Print(Name)
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into shedule (name, type, date,time) values ($1,$2,$3,$4)", Name, Typ, Datum, Vremya)
	if err != nil {
		panic(err)
	}

	log.Print(result)

}
