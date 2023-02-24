package http

import "fmt"

const (
	paramItemID = "itemID"
)

var (
	PathGetItem    = fmt.Sprintf("/items/:%s", paramItemID)
	PathSaveItem   = "/items"
	PathUpdateItem = fmt.Sprintf("/items/:%s", paramItemID)
)
