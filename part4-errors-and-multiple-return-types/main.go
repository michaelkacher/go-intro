package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type user struct {
	name  string
	email string
}

type contractor struct {
	user
	Rate        int
	HoursWorked int
}

// This is a pattern you will see in the Go core libraries.
// Note how this is capitalized, that means it can be imported and consumed
// by another package
var NegativeSalaryError = errors.New("yarg, must be > 0")

// this has two return values
// the signature could also be written func (c contractor) getSalary() (int, error)
// if you do this, return 0, NegativeSalaryError
func (c contractor) getSalary() (salary int, err error) {
	if c.Rate <= 0 || c.HoursWorked <= 0 {
		salary = 0
		err = NegativeSalaryError
		return // could be return 0, err
	}

	salary = c.Rate * c.HoursWorked

	return // could be return salary, nil
}

func main() {
	c := contractor{user{"contractorBob", "myemail@domain.com"}, 45, -50}
	salary, err := c.getSalary()

	fmt.Println(salary)
	fmt.Println(err.Error())

	// errors are handled like this throughout go.  Here is a json example, checking
	// that there is no error
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error serializing contractor")
	}

	// the 'string()' is casing the value of b to a string
	fmt.Println(string(b))

}
