package main

import(
	"fmt"
)
//var email strin = "Opitakorn_k@su.ac.th"
func main(){
	// var name string = "kanisorn"
	var age int = 20
	email := "Opitakorn_k@su.ac.th"
	gpa := 3.69
	firstname, lastname := "kanisorn", "opitakorn"
	fmt.Printf("Name %s %s, Age %d, Email %s, Gpa %.2f\n",firstname,lastname,age,email,gpa)
}