package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, v := range args {
		if v == "galaxy" || v == "galaxy 01" || v == "01" {
			fmt.Println("Alert!!!")

			return
		}
	}
}
