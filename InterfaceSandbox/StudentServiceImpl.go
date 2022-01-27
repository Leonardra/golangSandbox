package main

import (
	"errors"
	"fmt"
)

type StudentServiceImpl struct{
	students map[uint64] Student
}


func createService() *StudentServiceImpl{
	studentService := new(StudentServiceImpl)
	studentService.students = make(map[uint64]Student)
	return studentService
}

func (studentService *StudentServiceImpl) addStudent(student *Student) (int, error){
	studentService.students = make(map[uint64]Student)
	if student == nil{
		return 0, errors.New("student cannot be null")
	}
	studentService.students[uint64(student.id)] = *student
	fmt.Println(len(studentService.students))
	return len(studentService.students), nil
}

func (studentService *StudentServiceImpl)deleteStudent(id uint64) int{
	for key, _ := range studentService.students{
		if key == id{
			delete(studentService.students, key)
		}
	}
	return len(studentService.students)
}
