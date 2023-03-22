package docker

type ListVersionsResponse struct {
	Service  string    `json:"service"`
	Versions []Version `json:"versions"`
}

type Version struct {
	ImageID    string `json:"image_id"`
	Tag        string `json:"tag"`
	Containers int64  `json:"containers"`
	Size       string `json:"size"`
	Created    string `json:"created"`
}
