package main

import (
	"fmt"

	"github.com/pavan-subhakar/fpl/users"
)

func main() {
	pavan := &users.User{}
	pavan.UpdateUser("pavan", "sample@example.com")
	fmt.Println(pavan.Email)
	fmt.Println(pavan.Name)
}
