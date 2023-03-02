package http

import "fmt"

const (
	paramItemID = "itemID"
)

var (
	PathMetrics    = "/metrics"
	PathGetItem    = fmt.Sprintf("/items/:%s", paramItemID)
	PathSaveItem   = "/items"
	PathUpdateItem = fmt.Sprintf("/items/:%s", paramItemID)
	PathDeleteItem = fmt.Sprintf("/items/:%s", paramItemID)
)
