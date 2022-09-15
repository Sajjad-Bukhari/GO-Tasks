// Sajjad Bukhari
// 19I-1686
// CySec-M
// Task # 03: Printing a list of students

package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollNo  int
	name    string
	address string
}

func NewStudent(rollNo int, name string, address string) *Student {
	s := new(Student)
	s.rollNo = rollNo
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	students []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	s := NewStudent(rollno, name, address)
	ls.students = append(ls.students, s)
	return s
}

// Write a function to print the details of all the students in student list

func (ls *StudentList) PrintStudentDetails() {
	counter := 1
	for _, s := range ls.students {

		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 25), counter, strings.Repeat("=", 25))
		fmt.Println("Roll No: ", s.rollNo)
		fmt.Println("Name: ", s.name)
		fmt.Println("Address: ", s.address)
		counter++
	}
}
func main() {
	ls := new(StudentList)
	ls.CreateStudent(1684, "Sajjad", "Lahore")
	ls.CreateStudent(1786, "Usama", "Karachi")
	ls.CreateStudent(2171, "Yaqoob", "Chitral")

	ls.PrintStudentDetails()
}