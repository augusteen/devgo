package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("num", 302, "number")

	flag.Parse()

	fmt.Println("word", *wordPtr)
	fmt.Println("numb", *numPtr)
}
