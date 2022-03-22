/*
	Print every rune code point of the uppercase alphabet three times.
	Your output should look like this:
		65
		U+0041 'A'
		U+0041 'A'
		U+0041 'A'
		66
		U+0042 'B'
		U+0042 'B'
		U+0042 'B'
		… through the rest of the alphabet characters
*/

package main

import "fmt"

func main() {
	for firstIndex := 65; firstIndex <= 90; firstIndex++ {
		fmt.Println(firstIndex)
		for secondIndex := 0; secondIndex < 3; secondIndex++ {
			fmt.Printf("%#U\n", firstIndex)
		}
	}
}
