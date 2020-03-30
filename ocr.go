package main

import (
  "fmt"
  "image"
  "image/png"
	"os"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	statement, err := os.Open("statement.png")
	if err != nil {
		fmt.printLn("")
	}
	client := gosseract.NewClient()

}
