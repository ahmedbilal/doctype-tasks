package main

import "fmt"

type Person struct {
	Age uint
}

func isTwiceOldAsSomeone(people []Person) bool {
	var peopleMap = make(map[uint]Person)

	for _, person := range people {
		peopleMap[person.Age] = person
		_, double_found := peopleMap[person.Age*2]

		// This default case caters for a person whose age is an odd number
		// because dividing an odd number by 2 results in a truncate operation
		var half_found = false

		if person.Age%2 == 0 {
			_, half_found = peopleMap[person.Age/2]
		}

		if double_found || half_found {
			return true
		}
	}
	return false
}

func isAtleastTwiceAsOldAsSomeone(people []Person) bool {
	// This function assumes that people array is not empty i.e it contain atleast one element
	var minMax = func(people []Person) (uint, uint) {
		var min = people[0].Age
		var max = people[0].Age
		for _, person := range people {
			if person.Age < min {
				min = person.Age
			}
			if person.Age > max {
				max = person.Age
			}
		}
		return min, max
	}
	var min, max = minMax(people)
	return max >= 2*min
}

func main() {
	var people = []Person{Person{1}, Person{3}}
	fmt.Println("isAtleastTwiceAsOldAsSomeone =", isAtleastTwiceAsOldAsSomeone(people))
	fmt.Println("isTwiceOldAsSomeone =", isTwiceOldAsSomeone(people))
}
