package main

import (
	"fmt"
	"strings"

	"github.com/Luiz-Wendel/go-studies/01-learn-how-to-code/12-testing-and-benchmarking/06-benchmark-examples/mystr"
)

const text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Justo nec ultrices dui sapien eget mi proin sed libero. Eget sit amet tellus cras adipiscing. Duis ultricies lacus sed turpis tincidunt id aliquet risus. Commodo ullamcorper a lacus vestibulum sed arcu non odio. Enim facilisis gravida neque convallis a cras semper auctor neque. Nisl suscipit adipiscing bibendum est. Nunc mi ipsum faucibus vitae aliquet nec ullamcorper sit amet. Urna neque viverra justo nec ultrices dui. Sit amet volutpat consequat mauris. In tellus integer feugiat scelerisque varius morbi. Sagittis eu volutpat odio facilisis mauris sit amet massa. Gravida neque convallis a cras semper auctor neque. Eget magna fermentum iaculis eu non diam phasellus vestibulum lorem."

func main() {
	words := strings.Split(text, " ")

	for _, word := range words {
		fmt.Println(word)
	}

	fmt.Printf("\n%s\n", mystr.Cat(words))
	fmt.Printf("\n%s\n", mystr.Join(words))
}
