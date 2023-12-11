package main

import "fmt"

type employee struct {
	employeeId string
	employeeName string
	phone string

}

func main() {
	// employee1 := employee{
	// 	employeeId: "101",
	// 	employeeName: "Auu",
	// 	phone: "0888888888",
	// }

	// fmt.Println("employee1 =", employee1)

	// employeeList := [3] employee {}
	// employeeList[0] = employee{
	// 	employeeId: "101",
	// 	employeeName: "Auu",
	// 	phone: "081",
	// }
	// employeeList[1] = employee{
	// 	employeeId: "102",
	// 	employeeName: "Ploy",
	// 	phone: "082",
	// }
	// employeeList[2] = employee{
	// 	employeeId: "103",
	// 	employeeName: "Wit",
	// 	phone: "083",
	// }

	// fmt.Println("employee =", employeeList)

	employeeList := [] employee {}
	employee1 := employee {
		employeeId: "101",
		employeeName: "Nim",
		phone: "",
	}
	employee2 := employee {
		employeeId: "102",
		employeeName: "Noi",
		phone: "",
	}
	employee3 := employee {
		employeeId: "103",
		employeeName: "Nook",
		phone: "",
	}

	employeeList = append(employeeList,employee1, employee2, employee3)

	fmt.Println("employee =", employeeList)

}