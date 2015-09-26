package main

import "fmt"

// Lower case 'user' means it cannot be consumed by another package
type user struct {
	name  string
	email string
}

// notice how user is embeded
// go does not have inheritance, you extend through composition
type contractor struct {
	user
	rate        int
	hoursWorked int
}

// creating another struct wither 'user' embedded
type employee struct {
	user
	rate int
}

// there are no classes in go
// you can invoke functionality on types with method receivers
// in this function, user is the method receiever, see it invoked below
func (u user) msg() {
	fmt.Printf("hello %s\n", u.name)
}

// interfaces in Go are implicit.  Any types with a method receiver that match
// the signature 'getSalary()' will be of this interface
type salaryable interface {
	getSalary() int
}

// type contractor is now salaryable, note that this happens at compile time
// singnature for functions:
// fun (method receiver) functionName(parameters...) (return values)
func (c contractor) getSalary() int {
	salary := c.rate * c.hoursWorked

	return salary
}

func (e employee) getSalary() int {
	salary := e.rate * 40

	return salary
}

func main() {
	// the := is implicitly determines type at compile time
	// this user in create with a composite literal, it looks like json
	// useful in constructors
	u := user{name: "Test User", email: "testuser@domain.com"}
	u.msg()

	c := contractor{user: user{name: "Contractor Cole", email: "Cole@domain.com"}, rate: 22, hoursWorked: 55}
	// notice how we can call the embedded user method 'msg()'
	c.user.msg()
	// this method is promoted up, we don't have to include the embedded user
	c.msg()

	// compsite literals can explicitly list the fields like above or be composed of just the values
	e := employee{user{"Employee Edward", "Edward@domain.com"}, 15}
	e.msg()

	// this is creating a slice of 'salaryable' items, I will cover slices more later
	people := make([]salaryable, 3)
	people[0] = employee{user{"Rob1", "Rob1@domain.com"}, 15}
	people[1] = employee{user{"Rob2", "Rob2@domain.com"}, 30}
	people[2] = contractor{user{"ContractorBob", "myemail@domain.com"}, 45, 50}

	// Look! polymorphic behavior without classes.  We were able to invoke getSalary()
	// on both employee and contractor
	// the '_' could have also been a variable that kept track of the index count
	// of the loop.  Putting the _ denotes that we don't care about the index for this
	// loop.
	for _, p := range people {
		salary := p.getSalary()
		fmt.Println(salary)
	}
}
