package docker

type ContainerStatus int

const (
	ContainerStatusRunning ContainerStatus = iota
)

// String returns the ContainerStatus description as string
func (containerStatus ContainerStatus) String() string {
	switch containerStatus {
	case ContainerStatusRunning:
		return "running"
	}
	return "unknown"
}

// ContainerStatusFrom parses a value as a ContainerStatus
func ContainerStatusFrom(value string) ContainerStatus {
	switch value {
	case "running":
		return ContainerStatusRunning
	}
	return ContainerStatus(-1)
}
