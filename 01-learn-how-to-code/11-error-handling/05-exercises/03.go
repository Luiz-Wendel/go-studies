/*
  Create a struct “customErr” which implements the builtin.error interface.
  Create a func “foo” that has a value of type error as a parameter.
  Create a value of type “customErr” and pass it into “foo”.
  If you need a hint, [here is one](https://play.golang.org/p/L5QWV8-p11).
*/

package main

import (
	"fmt"
	"log"
)

type customErr struct {
	message string
}

func (ce *customErr) Error() string {
	return fmt.Sprintf("Error: %v", ce.message)
}

func main() {
	ce := customErr{
		message: "xablau",
	}

	foo(&ce)
}

func foo(ce error) {
	log.Fatalln(ce.(*customErr).message)
}
