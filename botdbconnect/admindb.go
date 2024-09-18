package botdbconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "95.174.94.157"
// 	port     = 5432
// 	user     = "asmolin"
// 	password = "oadnpvia"
// 	dbname   = "bot"
// )

type UsersData struct {
	name    string
	balance string
}

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func GetUsersFromLastEvent() string {

	var users string
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select name from users")
	userdata := []UsersData{}

	for rows.Next() {
		p := UsersData{}
		err := rows.Scan(&p.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		userdata = append(userdata, p)
	}
	for _, p := range userdata {
		users := fmt.Sprint(p.name)
		return users
	}
	return users
}
