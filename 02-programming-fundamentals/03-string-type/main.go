package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// converts string to slice of byte
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		// prints utf-8 code point
		fmt.Printf("%#U\n", s[i])
	}

	for i, v := range s {
		// prints index and decimal value
		fmt.Println(i, v)

		// prints hexadecimal value
		fmt.Printf("hex value: %#x\n", v)
	}
}
