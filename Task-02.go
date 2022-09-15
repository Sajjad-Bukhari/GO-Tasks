// Sajjad Bukhari
// 19I-1686
// CySec-M
// Task # 02: Array of Structs

package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func main() {
	var emp1 employee
	var emp2 employee
	var emp3 employee
	var comp company

	// emp1 Specifications
	emp1.name = "Aftab"
	emp1.salary = 3000
	emp1.position = "Rector"

	// emp2 Specifications
	emp2.name = "Waseem"
	emp2.salary = 2500
	emp2.position = "Director"

	// emp3 Specifications
	emp3.name = "Kashif"
	emp3.salary = 2000
	emp3.position = "In-Charge"

	emplys := []employee{emp1, emp2, emp3}

	// comp Specifications
	comp.companyName = "FAST NUCES"
	comp.employees = emplys

	// Printing company details
	fmt.Println("Name: ", comp.companyName)
	fmt.Println("Employees: ", comp.employees)
}