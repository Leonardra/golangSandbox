package main


type StudentService interface{
	addStudent(student *Student) int
	deleteStudent(id uint64) int
}
