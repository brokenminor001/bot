package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "95.174.94.157"
	port     = 5432
	user     = "asmolin"
	password = "oadnpvia"
	dbname   = "bot"
)

type Shedule struct {
	name string
	date string
	time string
}

func GetSheduleTraining() {
	currentdate := fmt.Sprintln(time.Now().Format("2006-01-02"))
	fmt.Print(currentdate)
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select name,date from shedule where date > $1", currentdate)
	getshedule := []Shedule{}
	for rows.Next() {
		p := Shedule{}
		err := rows.Scan(&p.name, &p.date)
		if err != nil {
			fmt.Println(err)
			continue
		}
		getshedule = append(getshedule, p)
	}
	for _, p := range getshedule {
		fmt.Println(p.name, p.date)
	}

}

func main() {
	GetSheduleTraining()

}
