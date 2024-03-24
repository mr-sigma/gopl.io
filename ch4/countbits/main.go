package main

import (
	"flag"
	"fmt"
	"os"
	"slices"

	"golang.org/x/crypto/sha3"
)

var flagSha384 = flag.Bool("sha384", false, "use sha384")
var flagSha512 = flag.Bool("sha512", false, "use sha512")

func main() {
	flag.Parse()
	flags := flag.Args()

	fmt.Println("os.Args")
	for i, val := range os.Args {
		fmt.Println(i, val)
	}

	fmt.Println("flags")
	for i, val := range flags {
		fmt.Println(i, val)
	}
	toHash := []byte(os.Args[1])

	if slices.Contains(flags, "flagSha384") {
		fmt.Println(sha3.Sum384(toHash))
	} else if slices.Contains(flags, "flagSha512") {
		fmt.Println(sha3.Sum512(toHash))
	} else {
		fmt.Println(sha3.Sum256(toHash))
	}
}
