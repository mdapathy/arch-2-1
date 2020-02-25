package main

import (
	"fmt"
	"github.com/mdapathy/arch-design-1"
)

func main() {
	fmt.Println(buildVersion)
	//Parse from command line TODO

	if res, err := lab1.PrefixEvaluation("+ 2 2"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
