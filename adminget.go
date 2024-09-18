package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var Res string

type SheduleData struct {
	name        string
	sheduletype string
	date        string
	time        string
}

const (
	host     = "95.174.94.157"
	port     = 5432
	user     = "asmolin"
	password = "oadnpvia"
	dbname   = "bot"
)

// Table of shedule Tree positions display
// Position Last
func GetLastSheduleOne() string {

	var last_shedule = []SheduleData{}
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name,type,date,time FROM (SELECT * FROM shedule ORDER BY id DESC LIMIT 1) as r ORDER BY id;")
	//rows, err := db.Query("SELECT type as sheduletype, date, time from shedule;")
	for rows.Next() {
		p := SheduleData{}
		err := rows.Scan(&p.name, &p.sheduletype, &p.date, &p.time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		last_shedule = append(last_shedule, p)
		fmt.Println(last_shedule)

	}
	for _, p := range last_shedule {
		fmt.Println(p.date)
		var Res string = fmt.Sprintln(p.name, p.sheduletype, p.date, p.time)
		return Res

	}
	return Res
}
func GetBeforeLastSheduleOne() string {

	var last_shedule = []SheduleData{}
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name,date,time FROM (SELECT * FROM shedule ORDER BY id DESC LIMIT 2) as r ORDER BY id;")

	for rows.Next() {
		p := SheduleData{}
		err := rows.Scan(&p.name, &p.date, &p.time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		last_shedule = append(last_shedule, p)
		fmt.Println(last_shedule)

	}
	for _, p := range last_shedule {
		fmt.Println(p.date)
		var Res string = fmt.Sprintln(p.name, p.date, p.time)
		return Res

	}
	return Res
}
func GetBeforeBeforeLastSheduleOne() string {

	var last_shedule = []SheduleData{}
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name,date,time FROM (SELECT * FROM shedule ORDER BY id DESC LIMIT 3) as r ORDER BY id;")

	for rows.Next() {
		p := SheduleData{}
		err := rows.Scan(&p.name, &p.date, &p.time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		last_shedule = append(last_shedule, p)
		fmt.Println(last_shedule)

	}
	for _, p := range last_shedule {
		fmt.Println(p.date)
		var Res string = fmt.Sprintln(p.name, p.date, p.time)
		return Res

	}
	return Res
}

func main() {
	var resultatlast string = GetLastSheduleOne()
	var resultatbeforelast string = GetBeforeLastSheduleOne()
	var resultatbeforebeforelast string = GetBeforeBeforeLastSheduleOne()
	fmt.Print(resultatbeforebeforelast)
	http.HandleFunc("/getlast", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Fprintf(w, resultatlast)

	})
	http.HandleFunc("/getbeforelast", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Fprintf(w, resultatbeforelast)

	})
	http.HandleFunc("/getbeforebeforelast", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Fprintf(w, resultatbeforebeforelast)

	})
	http.ListenAndServe(":8082", nil)

}
