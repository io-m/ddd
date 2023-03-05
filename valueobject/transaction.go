package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a value_object cause it has no ID and it is immutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
