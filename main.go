package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("Port:",portString)
}