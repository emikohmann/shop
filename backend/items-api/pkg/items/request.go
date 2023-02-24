package items

type GetItemRequest struct {
	ID int64
}

type SaveItemRequest struct {
	Item Item
}

type UpdateItemRequest struct {
	Item Item
}
