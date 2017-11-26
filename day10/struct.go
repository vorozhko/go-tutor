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
	var e Employee
	e.ID = 1
	e.Name = "Petia Pyatochkin"
	PrintEmployee(e)

	// Struct can reference on it's own type
	//Lets define manager for an employee
	var manager Employee
	manager.ID = 2
	manager.Name = "Middle Level"
	e.Manager = &manager
	PrintEmployee(e)

	var cto Employee
	cto.ID = 3
	cto.Name = "CTO"
	manager.Manager = &cto
	//should print 3 level org structure
	PrintEmployee(e)
}

// PrintEmployee - print data in nice format
func PrintEmployee(e Employee) {
	fmt.Printf("ID = %d\nName = %s\n", e.ID, e.Name)
	//if ID is set than employee has a manager
	//other way is to compare e.Manager to Employee{} type like if (Employee{}) == e.Manager  {
	if e.Manager != nil {
		fmt.Printf("Manager of %s:\n", e.Name)
		PrintEmployee(*e.Manager)
		return
	}
	fmt.Print("----------\n\n")
}
