package main

import "fmt"


func main() {
	var courseName [] string
	courseName = []string {"Math", "Sciece", "Art"}
	// fmt.Println(courseName)
	
	courseName = append(courseName, "Thai", "English")
	fmt.Println(courseName)
	
	courseWeb := courseName[0:2] 
	// fmt.Println(courseWeb)

	courseWeb = courseName[1:] 
	fmt.Println(courseWeb)
}