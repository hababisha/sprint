package initializr

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	Name         string
	BootVersion  string
	JavaVersion  string
	Build        string
	Packaging    string
	Dependencies []string
}

type Client struct {
	BaseURL string
	HTTP    *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: "https://start.spring.io",
		HTTP:    http.DefaultClient,
	}
}

func (c *Client) Generate(config Config) (*http.Response, error) {
	params := url.Values{}

	params.Set("type", config.Build+"-project")
	params.Set("language", "java")
	params.Set("bootVersion", config.BootVersion)
	params.Set("packaging", config.Packaging)
	params.Set("javaVersion", config.JavaVersion)
	params.Set("name", config.Name)

	params.Set(
		"dependencies",
		strings.Join(config.Dependencies, ","),
	)

	req, err := http.NewRequest(
		http.MethodGet,
		c.BaseURL+"/starter.zip?"+params.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	return c.HTTP.Do(req)
}

func Save(response *http.Response, path string) error {
	defer response.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)

	return err
}
