package main

import "fmt"
import "rsc.io/quote"
import "morestrings/morestrings.go"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}