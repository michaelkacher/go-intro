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

func main() {
	// the := is implicitly determines type at compile time
	// this user in create with a composite literal, it looks like json
	// useful in constructors
	u := user{name: "Test User", email: "testuser@domain.com"}
	u.msg()

	contractor := contractor{user: user{name: "Contractor Cole", email: "Cole@domain.com"}, rate: 22, hoursWorked: 55}
	// notice how we can call the embedded user method 'msg()'
	contractor.user.msg()
	// this method is promoted up, we don't have to include the embedded user
	contractor.msg()

	// compsite literals can explicitly list the fields like above or be composed of just the values
	e := employee{user{"Employee Edward", "Edward@domain.com"}, 15}
	e.msg()
}
