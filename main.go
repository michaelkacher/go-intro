package main

import "fmt"

// Lower case 'user' means it cannot be consumed by another package
type user struct {
	name  string
	email string
}

// a plain function
func hello() {
	fmt.Println("hello world")
}

// there are no classes in go
// you can invoke functionality on types with method receivers
// in this function, user is the method receiever, see it invoked below
func (u user) msg() {
	fmt.Printf("hello %s\n", u.name)
}

func main() {
	hello()

	var u1 user
	u1.name = "Bob"
	u1.msg()

	// the := is implicitly determines type at compile time
	// this user in create with a composite literal, it looks like json
	// useful in constructors
	u2 := user{name: "Test User", email: "testuser@domain.com"}
	u2.msg()
}
