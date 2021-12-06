package main

import (
	"fmt"

	"github.com/AtomicMegaNerd/rcd_go_types/organization"
)

func main() {
	p := organization.NewPerson("Chris", "Dunphy")
	err := p.SetTwitterHandler("@atomicmeganerd")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectUrl())
	fmt.Println(p.FullName())
	fmt.Println(p.ID())
}
