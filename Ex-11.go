// Sajjad Bukhari
// 19I-1686
// CySec-M
// Example # 11

package main

import "fmt"

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "Bahria Town")
	student.CreateStudent(25, "Naveed", "DHA")
	fmt.Println(student)
}