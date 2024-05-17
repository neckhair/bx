package internal

import (
	"io"
	"net/http"

	"github.com/neckhair/bx/bexio"
)

type Client interface {
	Get(url string, query map[string]string) (*http.Response, error)
	Post(url string, query map[string]string, body io.Reader) (*http.Response, error)
	BaseUrl() string

	ListContacts(limit int) ([]bexio.Contact, error)

	ListProjects(limit int) (bexio.ProjectList, error)
	GetProjectsByState(state bexio.ProjectState) (bexio.ProjectList, error)
}
