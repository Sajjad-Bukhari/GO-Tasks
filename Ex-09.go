// Sajjad Bukhari
// 19I-1686
// CySec-M
// Example # 09

package main

import "fmt"

type StudentRecord struct {
	rollnumber int
	name       string
	address    string
}

func AddStudent(rollno int, name string, address string) *StudentRecord {
	s := new(StudentRecord)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

func main() {
	st := AddStudent(24, "Ehtesham", "aaaaaaaaaa")
	fmt.Println(&st, st)

	st = AddStudent(25, "Naveed", "bbbbbbbbbb")
	fmt.Println(&st, st)
}