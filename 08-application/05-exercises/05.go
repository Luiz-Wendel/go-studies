/*
	Starting with [this code](https://play.golang.org/p/BVRZTdlUZ_), sort the []user by
		● age
		● last
	Also sort each []string “Sayings” for each user
		● print everything in a way that is pleasant
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byLast []user

func (l byLast) Len() int           { return len(l) }
func (l byLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l byLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func prettyPrint(users []user) {
	for _, user := range users {
		fmt.Println("\n", user.First, user.Last, user.Age)

		sort.Strings(user.Sayings)

		for _, saying := range user.Sayings {
			fmt.Println("\t", saying)
		}
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	prettyPrint(users)

	fmt.Println("\n-----")

	sort.Sort(byAge(users))
	prettyPrint(users)

	fmt.Println("\n-----")

	sort.Sort(byLast(users))
	prettyPrint(users)
}
