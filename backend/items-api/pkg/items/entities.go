package items

import "time"

type ItemList struct {
	Paging Paging
	Items  []Item
}

type Paging struct {
	Total  int
	Limit  int
	Offset int
}

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
	Punctuation  int
	DateCreated  time.Time
	LastUpdated  time.Time
}
