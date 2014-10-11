package main

import (
	"github.com/jmrobles/outbarriers"
	"log"
)

func main() {

	log.Printf("Outbarriers Server")
	ctx := outbarriers.NewContext()
	ctx.Init()
	err := ctx.StartDB()
	if err != nil {
		log.Fatal("Can't connect to DB")
	}
	ctx.Start()

}
