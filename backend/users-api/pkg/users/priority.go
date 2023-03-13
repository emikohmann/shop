package users

type Priority int

const (
    PriorityLow Priority = iota
    PriorityMedium
    PriorityHigh
)

// Value returns the priority numeric value
func (priority Priority) Value() uint8 {
    switch priority {
    case PriorityLow:
        return 1
    case PriorityMedium:
        return 2
    case PriorityHigh:
        return 3
    }
    return 0
}
