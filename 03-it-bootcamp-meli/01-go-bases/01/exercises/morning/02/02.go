package main

import "fmt"

func main() {
	temperature := 12.8
	umidity := 38
	atmosfericPressure := 30.09

	fmt.Printf("Temperature\t\t%v\t%T\n", temperature, temperature)
	fmt.Printf("Umidity\t\t\t%v\t%T\n", umidity, umidity)
	fmt.Printf("Atmosferic pressure\t%v\t%T\n", atmosfericPressure, atmosfericPressure)
}
