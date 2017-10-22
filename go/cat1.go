package main

import "fmt"
import "flag"

func main() {
	catname := flag.String("n", " ", "hhhh")
	flag.Parse()
	fmt.Printf("catname", *catname)

}
