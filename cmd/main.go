package main

import (
	"log"

	"github.com/pr194/Collaborative-tool/cmd/server"

	"fmt"
)

func main() {

	srv := server.NewServer()
	if err := srv.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("🚀 Server is starting on :8080...")
	srv.Start("8080")

}
