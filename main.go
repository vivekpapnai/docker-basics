package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"docker-example/server"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Connected")
	srv := server.SetupRoutes()
	err := srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
