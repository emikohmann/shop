package items

type GetItemRequest struct {
	ID int64
}

type ListItemsRequest struct {
	Limit  int
	Offset int
}

type SaveItemRequest struct {
	Item Item
}

type UpdateItemRequest struct {
	Item Item
}

type DeleteItemRequest struct {
	ID int64
}
