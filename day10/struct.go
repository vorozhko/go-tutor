package main

import "fmt"

// Employee - new type
type Employee struct {
	ID      int
	Name    string
	Manager *Employee
}

func main() {
	// Example of initializing new struct data
	var worker Employee
	worker.ID = 1
	worker.Name = "Petia Pyatochkin"
	PrintEmployee(worker)

	// Struct can reference on it's own type
	//Lets define manager for an employee
	var manager Employee
	manager.ID = 2
	manager.Name = "Middle Level"
	//using pointer we create a reference to manager struct and keep hierachy of oranization
	worker.Manager = &manager
	PrintEmployee(worker)

	//define fields ID and Name using struct literals
	var cto = Employee{ID: 3, Name: "cto"}
	manager.Manager = &cto
	//should print 3 level org structure
	PrintEmployee(worker)
}

// PrintEmployee - print data in nice format
func PrintEmployee(e Employee) {
	fmt.Printf("ID = %d\nName = %s\n", e.ID, e.Name)
	//if ID is set than employee has a manager
	//other way is to compare e.Manager to Employee{} type like if (Employee{}) == e.Manager  {
	if e.Manager != nil {
		fmt.Printf("Manager of %s:\n", e.Name)
		//recursively go through managers tree
		PrintEmployee(*e.Manager)
		return
	}
	fmt.Print("----------\n\n")
}
