package main

import (
	"KreditBee-assement/apicalls"
	"KreditBee-assement/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	r := router.Router()
	apicalls.InserttoDatabase()
	fmt.Println("Insert done!!")
	fmt.Println("Starting server on the port 8080...")

    log.Fatal(http.ListenAndServe(":8080", r))

}