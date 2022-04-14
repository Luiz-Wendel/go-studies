/*
	Starting with [this code](https://play.golang.org/p/b_UuCcZag9), unmarshal the JSON into a Go data structure.

	Hint: https://mholt.github.io/json-to-go/

	● code setup (just fyi, not needed for exercise):
		○ https://play.golang.org/p/nWPP5Z2Q4e
*/

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	bs := []byte(s)

	var people []person

	error := json.Unmarshal(bs, &people)

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("\n-----")

		fmt.Println("\n", people)

		fmt.Println("\n-----")

		for _, person := range people {
			fmt.Println("\n", person.First, person.Last, person.Age)
			for _, saying := range person.Sayings {
				fmt.Println("\t", saying)
			}
		}
	}
}
