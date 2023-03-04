package http

import "fmt"

const (
	paramItemID = "itemID"
)

const (
	queryLimit  = "limit"
	queryOffset = "offset"
)

var (
	PathDocs       = "/docs/*any"
	PathMetrics    = "/metrics"
	PathGetItem    = fmt.Sprintf("/items/:%s", paramItemID)
	PathListItems  = "/items"
	PathSaveItem   = "/items"
	PathUpdateItem = fmt.Sprintf("/items/:%s", paramItemID)
	PathDeleteItem = fmt.Sprintf("/items/:%s", paramItemID)
)
