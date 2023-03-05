package main

import (
	"log"

	"github.com/io-m/ddd/aggregate"
)

func main() {
	customer, err := aggregate.NewCustomer("JOhn")
	if err != nil {
		log.Printf("Error %v", err)
		return
	}

	log.Printf("Customer here %v", customer)
}
