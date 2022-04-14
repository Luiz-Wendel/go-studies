/*
  Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
*/

package main

import "fmt"

func main() {
	for index := 10; index <= 100; index++ {
		fmt.Printf("%v divided by 4, remains: %v\n", index, index%4)
	}
}
