package main

import (
	"context"
	"github.com/rmbziiik/Golang-application-development/Homework#6/internal/http"
	"github.com/rmbziiik/Golang-application-development/Homework#6/internal/store/inmemory"
	"log"
)

func main() {
	store := inmemory.NewDB()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()
}