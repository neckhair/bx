package bexio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

type ProjectList []Project

func parseListProjectsResponse(body io.Reader) (ProjectList, error) {
	var list ProjectList

	projects, err := io.ReadAll(body)
	if err != nil {
		return ProjectList{}, errors.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(projects, &list)
	if err != nil {
		return ProjectList{}, errors.Errorf("could not parse project list: %v", err)
	}

	return list, nil
}

func (c *Client) ListProjects(limit int) (ProjectList, error) {
	projectsUrl := fmt.Sprintf("%s/pr_project", c.BaseUrl)
	params := QueryParams{"limit": strconv.Itoa(limit)}
	resp, err := c.Get(projectsUrl, params)
	if err != nil {
		return nil, err
	}

	return parseListProjectsResponse(resp.Body)
}

func (c *Client) GetProjectsByState(state ProjectState) (ProjectList, error) {
	criteria := []SearchCriteria{NewSearchCriteria("pr_state_id", "=", fmt.Sprintf("%d", state))}
	b, err := json.Marshal(criteria)
	if err != nil {
		return nil, err
	}

	projectsUrl := fmt.Sprintf("%s/pr_project/search", c.BaseUrl)
	params := QueryParams{}
	resp, err := c.Post(projectsUrl, params, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	return parseListProjectsResponse(resp.Body)
}
