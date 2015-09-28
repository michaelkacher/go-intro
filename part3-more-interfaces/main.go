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

type reporter interface {
	salaryable
	notify(data int)
}

func (u user) notify(data int) {
	fmt.Printf("Sending %d email to: %s\n", data, u.email)
}

func generateReport(r reporter) {
	reportData := r.getSalary()

	r.notify(reportData)
}

func main() {
	reportable := make([]reporter, 3)
	reportable[0] = employee{user{"Rob1", "Rob1@domain.com"}, 15}
	reportable[1] = employee{user{"Rob2", "Rob2@domain.com"}, 30}
	reportable[2] = contractor{user{"ContractorBob", "myemail@domain.com"}, 45, 50}

	for _, r := range reportable {
		generateReport(r)
	}
}
