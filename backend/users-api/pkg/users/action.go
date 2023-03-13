package users

type Action int

const (
    ActionGet Action = iota
    ActionList
    ActionSave
    ActionUpdate
    ActionDelete
)

// String returns the action name
func (action Action) String() string {
    switch action {
    case ActionGet:
        return "GET"
    case ActionList:
        return "LIST"
    case ActionSave:
        return "SAVE"
    case ActionUpdate:
        return "UPDATE"
    case ActionDelete:
        return "DELETE"
    }
    return "UNKNOWN"
}
