// Package dog allows us to more fully understand dogs.
package dog

// Years converts human years to dog years.
func Years(age int) int {
	return age * 7
}

// YearsTwo converts human years to dog years.
func YearsTwo(age int) int {
	count := 0

	for i := 0; i < age; i++ {
		count += 7
	}

	return count
}
