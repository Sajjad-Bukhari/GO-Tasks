// Sajjad Bukhari
// 19I-1686
// CySec-M
// Task # 01: Pass Struct as Function Arguments

package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func DisplayPers(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary, "\n")
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 Specifications
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 Specifications
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Passing struct to DisplayPers function
	DisplayPers(pers1)
	DisplayPers(pers2)
}