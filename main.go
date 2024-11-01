package main

import (
	"net/http"
	"workout/infrastructure/routers"
)

func main() {
	router := routers.RootRouter()
	http.ListenAndServe(":3000", router)
}
