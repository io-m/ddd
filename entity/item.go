// Entity package holds all entites that are shared across subdomains
package entity

import "github.com/google/uuid"

type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
