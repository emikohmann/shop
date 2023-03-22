package docker

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Handlers struct {
	Client *Client
}

func NewHandlers(client *Client) *Handlers {
	return &Handlers{
		Client: client,
	}
}

func (handlers Handlers) CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}

func (handlers Handlers) ServiceVersions(c *gin.Context) {
	service := c.Query("service")

	images, err := handlers.Client.ListImages()
	if err != nil {
		fmt.Println("internal error", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	versions := make([]Version, 0)
	for _, image := range images {
		for _, tag := range image.RepoTags {
			components := strings.Split(tag, ":")
			if len(components) > 1 && components[0] == service {
				versions = append(versions, Version{
					ImageID:    image.ID,
					Tag:        components[1],
					Containers: image.Containers,
					Size:       ByteCountSI(image.Size),
					Created:    time.Unix(image.Created, 0).Format(time.RFC850),
				})
			}
		}
	}

	c.JSON(http.StatusOK, ListVersionsResponse{
		Service:  service,
		Versions: versions,
	})
}

func (handlers Handlers) ListContainers(c *gin.Context) {
	containers, err := handlers.Client.ListContainers()
	if err != nil {
		fmt.Println("internal error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"containers": containers,
	})
}

func (handlers Handlers) Build(c *gin.Context) {
	var request BuildRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("invalid request", err)
		c.Status(http.StatusBadRequest)
		return
	}
	if err := handlers.Client.Build(request); err != nil {
		fmt.Println("internal error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (handlers Handlers) Start(c *gin.Context) {
	var request StartRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("invalid request", err)
		c.Status(http.StatusBadRequest)
		return
	}
	if err := handlers.Client.Start(request); err != nil {
		fmt.Println("internal error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (handlers Handlers) Stop(c *gin.Context) {
	var request StopRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("invalid request", err)
		c.Status(http.StatusBadRequest)
		return
	}
	if err := handlers.Client.Stop(request); err != nil {
		fmt.Println("internal error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
