package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

// This is a type alias... it really is just a string
// type TwitterHandler = string

// This will not work with a type alias because you cannot add
// methods to string which is a non-local (not in this package) built-in type.
// func (th TwitterHandler) RedirectUrl() string {
// }

// This is a type definition
type TwitterHandler string

// Now we can extend our new type
func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

// Type definition
type socialSecurityNumber string

func NewSocialSecurityNumber(id string) Citizen {
	return socialSecurityNumber(id)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "Canada"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{id: id, country: country}
}

func (eur europeanUnionIdentifier) ID() string {
	return string(eur.id)
}

func (eur europeanUnionIdentifier) Country() string {
	return string(eur.country)
}

type Person struct {
	Name           // This is an embedded type, this is actually pretty cool.
	Citizen        // Embedding an interface in the type
	twitterHandler TwitterHandler
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{Name: Name{first: firstName, last: lastName},
		Citizen: citizen,
	}
}

func (p *Person) ID() string {
	// Here we are qualifying which identifier we are referring to as to
	// avoid conflict.  We are using the embedded Identifiable type.
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("Twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
