package main

import (
	"fmt"
	"log"
	"net/http"

	"webapp"
)

func gethandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Handler Executed")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Handler Executed")
}

func main() {

	wp := &webapp.WebApp{
		Router: webapp.NewRouter(),
	}

	wp.GET("/get", gethandler)
	wp.POST("/post", postHandler)

	s := &http.Server{
		Addr:    ":3000",
		Handler: wp,
	}

	log.Fatal(s.ListenAndServe())

}
