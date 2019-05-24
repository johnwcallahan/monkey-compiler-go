package main

import (
	"fmt"

	"github.com/johnwcallahan/monkey-go/token"
)

func main() {
	fmt.Println(token.LookupIdent("fn"))
}
