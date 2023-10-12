package main

import (
	"fmt"
)

type Employee struct {
	id, age             int
	firstName, lastName string
	smoker              bool
}

type Operation interface {
	createEmployee(*Employee)
	deleteEmployee(int) error
	updateEmplyee(*Employee)
	getEmployeeById(int) (*Employee, error)
}

type Database struct {
	employeeList []Employee
}

func (d *Database) initOneFakeData() {
	fmt.Println("---------intializing DB creating  table for employee----------")

	d.employeeList = append(d.employeeList, Employee{
		id:        1,
		age:       24,
		firstName: "Syed",
		lastName:  "Faran ul haq",
		smoker:    false,
	})
}
func (d *Database) createEmployee(newEmployee *Employee) {
	fmt.Println("---------new employee added----------")
	d.employeeList = append(d.employeeList, *newEmployee)
}
func (d *Database) deleteEmployee(indexToDelete int) error {
	fmt.Printf("---------deleteEmployee at element %v----------", indexToDelete)

	if indexToDelete >= 0 && indexToDelete < len(d.employeeList) {
		// Delete the element at the specified index
		d.employeeList = append(d.employeeList[:indexToDelete], d.employeeList[indexToDelete+1:]...)
		fmt.Println("After deletion:", d.employeeList)
		return nil
	} else {
		fmt.Println("Index out of bounds.")
		return fmt.Errorf("Index out of bounds.")
	}
}
func (d *Database) getEmployeeById(index int) (*Employee, error) {
	if index >= 0 && index < len(d.employeeList) {

		return &d.employeeList[index], nil
	} else {
		fmt.Println("Index out of bounds.")
		return nil, fmt.Errorf("Index out of bounds.")
	}
}

func main() {
	db := Database{employeeList: []Employee{}}
	db.initOneFakeData()
	fmt.Println(db.employeeList)
	newEmployee := Employee{1023, 40, "David", "Smith", true}
	db.createEmployee(&newEmployee)
	// fmt.Println(db.employeeList)
	// error := db.deleteEmployee(0)
	// if error != nil {
	// 	panic(error)
	// }
	emp, err := db.getEmployeeById(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(emp.lastName)
	fmt.Println(db.employeeList)
	//fmt.Printf("%+v\n", db)
}
