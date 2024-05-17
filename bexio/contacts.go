package bexio

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

func parseListContactsResponse(b []byte) ([]Contact, error) {
	var contacts []Contact

	err := json.Unmarshal(b, &contacts)
	if err != nil {
		return []Contact{}, errors.Errorf("could not parse contact list: %v", err)
	}

	return contacts, nil
}

func (c *Client) ListContacts(limit int) ([]Contact, error) {
	contactsUrl := fmt.Sprintf("%s/contact", c.BaseUrl)
	params := QueryParams{"limit": strconv.Itoa(limit)}
	resp, err := c.Get(contactsUrl, params)
	if err != nil {
		return nil, err
	}

	contacts, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Contact{}, errors.Errorf("error reading response body: %v", err)
	}

	return parseListContactsResponse(contacts)
}
