package organization

import (
	"errors"
	"fmt"
	"strings"
)

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

type Person struct {
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

func NewPerson(firstName, lastName string) Person {
	return Person{firstName: firstName, lastName: lastName}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

// Go implicitly implements interfaces if it has the
// correct method!
func (p *Person) ID() string {
	return "12345"
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
