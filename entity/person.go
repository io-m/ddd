// Entity package holds all entites that are shared across subdomains
package entity

import "github.com/google/uuid"

type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
