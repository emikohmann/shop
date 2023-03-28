package services

type Health int

const (
	Healthy Health = iota
	Unhealthy
)

// String returns the service health description as string
func (health Health) String() string {
	switch health {
	case Healthy:
		return "HEALTHY"
	case Unhealthy:
		return "UNHEALTHY"
	}
	return "UNKNOWN"
}
