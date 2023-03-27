package main

import (
	"FitaTechnicalTest-Golang/api"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err load env : ", err.Error())
	}
	api.Init()
}
