package main

import (
	"fmt"
	"github.com/mdapathy/arch-design-1"
	"log"
	"os"
)

func main() {


	fmt.Println("Build version ", buildVersion)

	if len(os.Args) < 2 {
		log.Fatalf("Not enough parameters specified, required 1, got %d", len(os.Args)-1)
	}

	if res, err := lab1.PrefixEvaluation(os.Args[1]); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
