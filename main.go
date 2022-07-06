package main

import (
	"strings"
	"fmt"
	"os"
)

func get_input_string() string {
	return os.Args[1]
}

func main() {
	s := get_input_string()
	fmt.Println(len(strings.Fields(s)))
}
