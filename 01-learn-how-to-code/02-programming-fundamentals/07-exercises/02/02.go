/*
	Using the following operators, write expressions and assign their values to variables:
		a. ==
		b. <=
		c. >=
		d. !=
		e. <
		f. >
	Now print each of the variables.
*/

package main

import "fmt"

func main() {
	a := 42 == 5
	b := 42 <= 5
	c := 42 >= 5
	d := "James Bond" != "Xablau"
	e := len("test") < len("thet")
	f := 42 > 5

	fmt.Println(a, b, c, d, e, f)
}
