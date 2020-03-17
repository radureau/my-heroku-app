package main

import (
	"github.com/radureau/my-heroku-app/herge-lab/server"
)

func main() {
	err := server.Serve()
	if err != nil {
		panic(err)
	}
}