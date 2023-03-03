package http

import "fmt"

const (
	paramItemID = "itemID"
)

var (
	PathDocs       = "/docs/*any"
	PathMetrics    = "/metrics"
	PathGetItem    = fmt.Sprintf("/items/:%s", paramItemID)
	PathSaveItem   = "/items"
	PathUpdateItem = fmt.Sprintf("/items/:%s", paramItemID)
	PathDeleteItem = fmt.Sprintf("/items/:%s", paramItemID)
)
