package botdbconnect

import (
	"database/sql"
	"fmt"
	"log"
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
	data string
}

type UserData struct {
	name    string
	balance string
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Checkuserifexist(t string) string {
	//t := strconv.Itoa(id)

	var database_id string
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select telegram_id from users where telegram_id=$1", t)
	CheckError(err)
	for rows.Next() {

		err = rows.Scan(&database_id)
		CheckError(err)

	}

	return database_id

}

func InsertNewUserID(userid string) {

	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into users (telegram_id) values ($1)",
		userid)
	if err != nil {
		panic(err)
	}
	log.Print(result)

}

func InsertNewUserName(name string, telegramid string) {

	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update users set name=$1 where telegram_id=$2",
		name, telegramid)
	if err != nil {
		panic(err)
	}
	log.Print(result)

}

func InsertNewUserSecondName(secondname string, telegramid string) {

	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update users set secondname=$1 where telegram_id=$2",
		secondname, telegramid)
	if err != nil {
		panic(err)
	}
	log.Print(result)

}

func SelectUserName(userid string) string {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select name from users where telegram_id = $1", userid)
	userdata := UserData{}
	err = row.Scan(&userdata.name)
	if err != nil {
		panic(err)
	}
	username := userdata.name
	return username
}
func GetBalance(userid string) string {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select balance from users where telegram_id = $1", userid)
	userbalance := UserData{}
	err = row.Scan(&userbalance.name)
	if err != nil {
		panic(err)
	}
	balance := userbalance.name
	return balance
}
func GetSheduleTraining() string {
	currentdate := fmt.Sprintln(time.Now().Format("2006-01-02"))
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select name,date,time from shedule where date > '$1'", currentdate)
	getshedule := Shedule{}
	err = row.Scan(&getshedule.data)
	if err != nil {
		panic(err)
	}
	raspisanie := getshedule.data
	return raspisanie

}

func Oplata(summa string, telegramid string) {

	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update users set balance=$1 where telegram_id=$2",
		summa, telegramid)
	if err != nil {
		panic(err)
	}
	log.Print(result)

}
func Zapis(telegramid string) {

	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into players(telegram_id) values($1)",
		telegramid)
	if err != nil {
		panic(err)
	}
	log.Print(result)

}
func GetSheduleForTemaOne() string {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select name from shedule where type='gameteamone' ORDER BY id DESC LIMIT 1")
	getshedule := Shedule{}
	err = row.Scan(&getshedule.data)
	if err != nil {
		panic(err)
	}
	raspisanie := getshedule.data
	return raspisanie

}
func GetSheduleForTemaTwo() string {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select name from shedule where type='gameteamtwo' ORDER BY id DESC LIMIT 1")
	getshedule := Shedule{}
	err = row.Scan(&getshedule.data)
	if err != nil {
		panic(err)
	}
	raspisanie := getshedule.data
	return raspisanie

}

func GetSheduleForTemaTree() string {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select name from shedule where type='gameteamtree' ORDER BY id DESC LIMIT 1")
	getshedule := Shedule{}
	err = row.Scan(&getshedule.data)
	if err != nil {
		panic(err)
	}
	raspisanie := getshedule.data
	return raspisanie

}

func GetSheduleForTemaFour() string {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select name from shedule where type='gameteamfour' ORDER BY id DESC LIMIT 1")
	getshedule := Shedule{}
	err = row.Scan(&getshedule.data)
	if err != nil {
		panic(err)
	}
	raspisanie := getshedule.data
	return raspisanie

}
