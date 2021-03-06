package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"data_scrapper/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	l := log.New(os.Stdout, "data_scrapper ", log.LstdFlags)
	postRouter := router.Methods(http.MethodPost).Subrouter()
	scrap := handlers.NewScrap(l)
	postRouter.HandleFunc("/scrap", scrap.ServerHTTP)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		l.Println("Starting server on port 9090")
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)

		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Shutting down ", sig)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)

}
