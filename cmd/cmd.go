package cmd

import (
	"fmt"
	"log"
	"net/http"
	"web/server"
)

func Cmd() {
	server := server.New()
	fmt.Printf("Starting server at port 8090\nhttp://localhost:8090/\n")
	if err := http.ListenAndServe(":8090", server.Handle()); err != nil {
		log.Fatal(err)
	}
}
