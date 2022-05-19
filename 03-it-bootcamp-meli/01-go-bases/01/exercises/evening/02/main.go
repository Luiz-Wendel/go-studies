package main

import "fmt"

type person struct {
	age, workingMonths, salary uint
	isEmployed                 bool
}

var (
	minimalAge    uint = 22
	isEmployed    bool = true
	workingMonths uint = 12
	salary        uint = 100000
)

func main() {
	p1 := person{age: 20, workingMonths: workingMonths, salary: salary, isEmployed: isEmployed}
	p2 := person{age: minimalAge, workingMonths: 11, salary: salary, isEmployed: isEmployed}
	p3 := person{age: minimalAge, workingMonths: workingMonths, salary: salary, isEmployed: false}
	p4 := person{age: minimalAge, workingMonths: workingMonths, salary: 6927, isEmployed: isEmployed}
	p5 := person{age: minimalAge, workingMonths: workingMonths, salary: salary, isEmployed: isEmployed}

	loans := []string{checkLoan(p1), checkLoan(p2), checkLoan(p3), checkLoan(p4), checkLoan(p5)}

	for _, loan := range loans {
		fmt.Println(loan)
	}
}

func checkLoan(p person) string {
	if p.age < minimalAge || p.isEmployed != isEmployed || p.workingMonths < workingMonths {
		return "Loan denied"
	}

	if p.salary < salary {
		return "Loan approved with interest"
	}

	return "Loan approved free of interest"
}
