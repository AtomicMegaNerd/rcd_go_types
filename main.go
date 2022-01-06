package main

import (
	"fmt"

	"github.com/AtomicMegaNerd/rcd_go_types/organization"
)

func main() {
	p := organization.NewPerson("Chris", "Dunphy", organization.NewSocialSecurityNumber("123-456-789"))
	err := p.SetTwitterHandler("@atomicmeganerd")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	fmt.Println("=================================")
	fmt.Println(p.FullName())
	fmt.Println(p.ID())
	fmt.Println(p.Country())
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectUrl())

	p2 := organization.NewPerson("Hans", "Pichert",
		organization.NewEuropeanUnionIdentifier("334-118-9920", "Czechia"),
	)
	fmt.Println("=================================")
	fmt.Println(p2.FullName())
	fmt.Println(p2.ID())
	fmt.Println(p2.Country())
}
