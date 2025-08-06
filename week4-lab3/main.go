package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID string `json:"id"`	
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	GPA float64 `json:"gpa"`
}
func(s*Student) IsHonor() bool {
	return s.GPA >= 3.5
}
func(s*Student) validate() error {
	if s.Name == "" {
		return errors.New("Name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("Year must be between 1-4")
	}
	if s.GPA < 0.0 || s.GPA > 4.0 {
		return errors.New("GPA must be between 0.0-4.0")
	}
	return nil
}

func main(){
	//var st Student = Student{ID:"1",Name:"",Email:"opitakorn_k@su.ac.th",Year:3,GPA:3.69}
	//st:= Student{ID:"1",Name:"kanisorn",Email:"opitakorn_k@su.ac.th",Year:3,GPA:3.69}
	students := []Student{
		{ID: "1", Name: "kanisorn", Email: "opitakorn_k@su.ac.th", Year: 3, GPA: 3.69},
		{ID: "2", Name: "john", Email: "john@example.com", Year: 4, GPA: 3.85},
	}
	newstudents := Student{ID: "3", Name: "jane", Email: "jane@example.com", Year: 2, GPA: 3.75}
	students = append(students, newstudents)
	
	for i, st := range students {
		fmt.Printf("%d Honor %v\n",i, st.IsHonor())
		fmt.Printf("%d Valid %v\n",i, st.validate())
	}
}