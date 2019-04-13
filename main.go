package main

import (
	"fmt"
	"net/http"

	"github.com/go-web-app/configs"
	"github.com/go-web-app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := configs.ConnectSql("localhost", "3306", "root", "password", "test")

	if err != nil {
		fmt.Println("DB error")
		//log.Fatal("Not able to connect DB")
		panic(err.Error())
	}
	RegisterHandlers(db)
}

func RegisterHandlers(db *configs.DB) {
	r := mux.NewRouter()
	handlers.RegisterBookHandlers(db, r)
	http.ListenAndServe(":8080", r)
}
