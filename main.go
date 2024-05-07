package main

import (
	"Api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Mongo db connected")
	r := router.Router()
	fmt.Println("server is getting started...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening at port 8080")
}
