package adminpost

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DataInputPost struct {
	Name string
}

func Poster() {

	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe(":8081", nil)

}

func PostHandler(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t DataInputPost
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Name)
	var tata = t.Name
	fmt.Print(tata)
}
