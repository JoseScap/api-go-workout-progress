package main

import (
	"go-workout-progress/infrastructure/routers"
	"net/http"
)

func main() {
	router := routers.RootRouter()
	http.ListenAndServe(":3000", router)
}
