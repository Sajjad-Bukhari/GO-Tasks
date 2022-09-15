// Sajjad Bukhari
// 19I-1686
// CySec-M
// Example # 10

package main

import "fmt"

type StudentRecord struct {
	rollnumber int
	name       string
	address    string
}

func (s *StudentRecord) AddStudent(rollno int, name string, address string) *StudentRecord {
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

func main() {
	st := new(StudentRecord)
	st.AddStudent(24, "Asim", "Bahria Town")
	fmt.Println(&st, st)
}