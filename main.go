package main

import (
	"fmt"

	"github.com/johnwcallahan/monkey-compiler-go/token"
)

func main() {
	fmt.Println(token.LookupIdent("fn"))
}
