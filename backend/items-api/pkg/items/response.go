package items

type GetItemResponse struct {
	Item Item
}

type ListItemsResponse struct {
	Paging Paging
	Items  []Item
}

type SaveItemResponse struct {
	Item Item
}

type UpdateItemResponse struct {
	Item Item
}

type DeleteItemResponse struct {
	ID int64
}
