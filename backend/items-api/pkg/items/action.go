package items

type Action int

const (
	ActionGet Action = iota
	ActionSave
	ActionUpdate
	ActionDelete
)

// String returns the action name
func (action Action) String() string {
	switch action {
	case ActionGet:
		return "GET"
	case ActionSave:
		return "SAVE"
	case ActionUpdate:
		return "UPDATE"
	case ActionDelete:
		return "DELETE"
	}
	return "UNKNOWN"
}
