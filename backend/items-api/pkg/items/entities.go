package items

import "time"

type Item struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	DateCreated time.Time
}
