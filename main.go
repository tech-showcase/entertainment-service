package main

import (
	"fmt"
	"github.com/tech-showcase/entertainment-service/cmd"
)

func main() {
	fmt.Println("Hi, I am Entertainment Service!")

	args := cmd.Parse()

	fmt.Println(args)
}
