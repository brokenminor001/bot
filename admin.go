package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "95.174.94.157"
	port     = 5432
	user     = "asmolin"
	password = "oadnpvia"
	dbname   = "bot"
)

type DataInputPost struct {
	PostRaspisanie string
}

func SendJqueryJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "script.js")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type UsersData struct {
	name    string
	balance string
}
type ViewData struct {
	Message string
	Title   string
}
type ViewDataDwa struct {
	Mess string
}
type DataInput struct {
	Name string
}

func GetUsersFromLastEvent() []UsersData {

	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select name,balance from users")
	userdata := []UsersData{}

	for rows.Next() {
		p := UsersData{}
		err := rows.Scan(&p.name, &p.balance)
		if err != nil {
			fmt.Println(err)
			continue
		}
		userdata = append(userdata, p)
	}
	return userdata

}
func main() {
	users := GetUsersFromLastEvent()

	http.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		data := ViewDataDwa{
			Mess: "",
		}
		tmpl, _ := template.ParseFiles("admin.html")
		tmpl.Execute(w, data)
		for _, p := range users {

			data := ViewData{
				Message: p.name,
				Title:   p.balance,
			}
			tmpl, _ := template.ParseFiles("spisok.html")
			tmpl.Execute(w, data)

		}

	})
	http.HandleFunc("/raspisanie", func(w http.ResponseWriter, r *http.Request) {
		data := ViewDataDwa{
			Mess: "",
		}
		tmpl, _ := template.ParseFiles("admin.html")
		tmpl.Execute(w, data)
		tmpl2, _ := template.ParseFiles("raspisaniye.html")
		tmpl2.Execute(w, data)

	})
	// http.HandleFunc("/postraspisanie", PostHandler)
	fmt.Println("Server is listening...")
	http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./script.js")
	})

	http.ListenAndServe(":80", nil)

}
