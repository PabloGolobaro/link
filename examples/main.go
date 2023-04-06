package main

import (
	"fmt"
	"github.com/pablogolobaro/link"
	"log"
	"os"
)

func main() {
	file, err := os.Open("ex3.html")
	if err != nil {
		log.Println(err)
		return
	}
	parse, err := link.Parse(file)
	if err != nil {
		return
	}
	fmt.Println(parse)
}
