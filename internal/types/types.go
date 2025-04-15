package types

import "time"

// Todo represents a task item.
type Todo struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}
