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
		log.Fatalf("Not enough parameters specified, required an expression, got %d parameters", len(os.Args)-1)
	}

	var input string

	for i := 1; i < len(os.Args); i++ {
		input += os.Args[i] + " "
	}

	if res, err := lab1.PrefixEvaluation(input); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
