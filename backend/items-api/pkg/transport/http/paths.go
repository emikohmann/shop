package http

import "fmt"

var (
	GetItem  = fmt.Sprintf("/items/:%s", paramItemID)
	SaveItem = "/items"
)
