package main

import (
	"fmt"
	"log"

	"crowdfunding/tools/seed"
)

func main() {
	if err := seed.Seed(); err != nil {
		log.Fatalf("seed failed: %v", err)
	}
	fmt.Println("seed finished")
}
