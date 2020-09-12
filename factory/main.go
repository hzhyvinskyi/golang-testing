package main

import (
	"fmt"
	"os"

	"github.com/hzhyvinskyi/ood-patterns/factory/tank_factory"
)

func main() {
	m1, err := tank_factory.Create("M1")
	if err != nil {
		fmt.Printf("Error occured: %s", err.Error())
		os.Exit(1)
	}
	m2, err := tank_factory.Create("M2")
	if err != nil {
		fmt.Printf("Error occured: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("M1: ", m1)
	fmt.Println("M2: ", m2)
	os.Exit(0)
}
