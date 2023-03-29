package static

import (
	"api/internal/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	fileLocation = "/pkg/static"
	fileName     = "static.json"
)

type Static struct {
	logger *logger.Logger
	data   Data
}

type Data struct {
	Services []Service `json:"services"`
}

type Service struct {
	Name      string    `json:"name"`
	Image     Image     `json:"image"`
	Container Container `json:"container"`
}

type Image struct {
	ID string `json:"id"`
}

type Container struct {
	ID   string `json:"id"`
	Port int    `json:"port"`
}

func NewStatic(logger *logger.Logger) (*Static, error) {
	static := &Static{
		logger: logger,
	}
	if err := static.Reload(); err != nil {
		return nil, err
	}
	return static, nil
}

// Reload reads file information and fulfill static config struct
func (static *Static) Reload() error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}
	location := fmt.Sprintf("%s%s/%s", workingDir, fileLocation, fileName)
	bytes, err := ioutil.ReadFile(location)
	if err != nil {
		return err
	}
	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	static.data = data
	return nil
}

// Data returns the current configuration
func (static *Static) Data() Data {
	return static.data
}
