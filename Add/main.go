package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"data_add/handlers"
	"os"
	"log"
	"time"
	"context"
	"os/signal"
)

func main() {
	router := mux.NewRouter()
	l := log.New(os.Stdout, "data_adder ", log.LstdFlags)
	postRouter := router.Methods(http.MethodPost).Subrouter()
	add := handlers.NewAdd(l)
	postRouter.HandleFunc("/add", add.ServerHTTP)
	
	server := &http.Server{
		Addr: ":9091",
		Handler: router,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	go func() {
		l.Println("Starting server on port 9091")
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)

		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	
	sig := <- sigChan
	l.Println("Shutting down ", sig)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
	
}