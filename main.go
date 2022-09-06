package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	scaledroneID     = "YOUR_SCALEDRONE_ID"
	scaledroneSecret = "YOUR_SCALEDRONE_SECRET"
	port             = ":8080"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/auth", auth).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static"))).Methods("GET")
	fmt.Printf("Server is runing on localhost", port)
	panic(http.ListenAndServe(port, r))

}
