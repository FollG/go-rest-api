package main

import (
	"go-rest-api/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("creating router")
	router := httprouter.New()

	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("starting application...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("server is listening port :8080")
	log.Fatalln(server.Serve(listener))
}
