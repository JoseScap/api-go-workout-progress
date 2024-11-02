package main

import (
	"net/http"
	"workout/infrastructure/database"
	"workout/infrastructure/routers"
)

func main() {
	database.ConnectDb()
	router := routers.RootRouter()
	http.ListenAndServe(":3000", router)
}
