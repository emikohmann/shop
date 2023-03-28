package services

type Status int

const (
	StatusNotRunning Status = iota
	StatusRunning
	StatusError
)

// String returns then service status description as string
func (status Status) String() string {
	switch status {
	case StatusNotRunning:
		return "NOT_RUNNING"
	case StatusRunning:
		return "RUNNING"
	case StatusError:
		return "ERROR"
	}
	return "UNKNOWN"
}
