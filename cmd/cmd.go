package main

import (
	"fmt"

	"github.com/homina/galaxy/pkg/input"
)

func main() {
	defer exception()
	input.CheckInput()
}

func exception() {
	if r := recover(); r != nil {
		fmt.Printf("%+v\n", r)
	}
}
