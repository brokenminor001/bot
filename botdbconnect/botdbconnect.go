package botdbconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "109.120.189.151"
	port     = 5432
	user     = "asmolin"
	password = "oadnpvia"
	dbname   = "bot"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkuserifexist(id string) string {
	var database_id string
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select user_id from users where user_id='$1'", id)
	CheckError(err)
	for rows.Next() {

		err = rows.Scan(&database_id)
		CheckError(err)

	}
	return database_id

}
