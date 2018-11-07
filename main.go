package main

import (
	"github.com/tommarler/Beard_Of_knowledge/system/app"
	"flag"
	"github.com/joho/godotenv"
    "log"
    "os"
)

var port string

func init() {

	flag.StringVar(&port, "port", "8000", "Assigning the port that the server should listen on.")

	flag.Parse()

	err := godotenv.Load("config.ini")
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}


func main() {

	s := app.NewServer()

	s.Init(port)
	s.Start()

}








