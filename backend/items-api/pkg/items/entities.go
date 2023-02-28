package items

import "time"

type Item struct {
	ID           int64
	Name         string
	Description  string
	Thumbnail    string
	Images       []string
	IsActive     bool
	Restrictions []string
	Price        float64
	Stock        int
	DateCreated  time.Time
	LastUpdated  time.Time
}
