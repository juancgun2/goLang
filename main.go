package main

import (
	"entregableGo/model"
	"fmt"
)

func main() {
	input := "TX02AB"

	result, err := model.NewResult(input)
	if err == nil {
		fmt.Print(result)
	} else {
		fmt.Println("Error detected: ")
		panic(err)
	}
}
