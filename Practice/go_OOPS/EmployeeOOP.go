package main

import "fmt"

var p1 = fmt.Println

//Can simply use this kind of global Map for simple storage, (Its alternate easy way)
var EmployeeDB = map[int]Employee{}

//Storing all the employee struct details in map
type AllEmployee struct {

	// EmployeeDB map[int]Employee

	EmployeeDB map[int]interface{} //Using interface, So we can store multiple structs Employee and ContractEmployee value in SaveEmployee() function
}

type Employee struct {
	Id         int
	Name       string
	Department Department
	Addresses  []*Address
	Mobile     []int
}

type ContractEmployee struct {
	Employee
	ThirdPartyClientPayrol string
}

type Department struct {
	DepartmentName     string
	DepartmentIncharge string
}

type Address struct {
	Street   string
	Locality string
	City     string
}

//Developed seperate struct pointer receiver methods for both AllEmployee and Employee structs

func (emp *Employee) NewEmployee(Name string, Id int) {

	emp.Id = Id
	emp.Name = Name
}

func (emp *Employee) setDepartment(Dname, Dincharge string) {
	department := Department{Dname, Dincharge}
	emp.Department = department
}

func (emp *Employee) setMobileNumber(Numbers []int) {
	emp.Mobile = Numbers
	p1("Added mobile number for the employee", emp.Name)
}

/*
func (employees *AllEmployee) SaveEmployee(emp Employee) { //(emp Employee) - Employee struct argument is not works for ContractEmployee struct, Since Golang not supports inheritance and only supports composition. So in this case We should use Interface, Check below example:
	employees.EmployeeDB[emp.Id] = emp
}
*/

func (employees *AllEmployee) SaveEmployee(emp interface{}) {

	//Interface type assertion with different struct types
	switch emp.(type) {

	case Employee:
		EmployeeObject := emp.(Employee)
		employees.EmployeeDB[EmployeeObject.Id] = EmployeeObject
	case ContractEmployee:
		ContractEmployeeObject := emp.(ContractEmployee)
		employees.EmployeeDB[ContractEmployeeObject.Id] = ContractEmployeeObject
	}

}

func (employees *AllEmployee) ViewAllEmployee() {
	p1("ALL Employees from Employees struct ----", employees)
}

func (employees *AllEmployee) deleteEmpoyee(employeeId int) {
	delete(EmployeeDB, employeeId)
	delete(employees.EmployeeDB, employeeId)
	p1("successfully deleted the employee", employeeId)
}

func EmployeeMainFunction() {

	var employees = AllEmployee{}
	employees.EmployeeDB = map[int]interface{}{} //Correct way of intializing the map with interface

	emp := Employee{}
	emp.NewEmployee("sathish", len(employees.EmployeeDB)+1)

	emp.setDepartment("Manufacturing", "Sundar")
	employees.SaveEmployee(emp)
	employees.ViewAllEmployee()

	emp.setMobileNumber([]int{9685475758, 9484635368})
	employees.SaveEmployee(emp) //Updating the employee values in different times and saving every time with "emp" object (or) We can save once finally after updated all the information for an employee.
	employees.ViewAllEmployee()

	emp = Employee{}
	emp.NewEmployee("Hari", len(employees.EmployeeDB)+1)
	emp.setDepartment("Admin", "Kali")
	employees.SaveEmployee(emp)
	employees.ViewAllEmployee()

	employees.deleteEmpoyee(emp.Id) //Removing one employee from the AllEmployee struct's map
	employees.ViewAllEmployee()

	contractEmp := ContractEmployee{}
	contractEmp.NewEmployee("Jai", len(employees.EmployeeDB)+1)
	contractEmp.setMobileNumber([]int{9685475758, 9484635369})
	employees.SaveEmployee(contractEmp)
	p1(contractEmp)
}

//Method overriding with the Employee struct's NewEmployee for adding custom values for contractEmployee struct
func (emp *ContractEmployee) NewEmployee(Name string, Id int) {

	emp.Id = Id

	//Adding below differents value for contract employee
	emp.Name = Name + "_ContractEmployee"
	emp.ThirdPartyClientPayrol = "crown Solutions Pvt Ltd"
}

func main() {

	EmployeeMainFunction()

}
