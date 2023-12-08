package bexio

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	ID        int       `json:"id"`
	Number    string    `json:"nr"`
	Name      string    `json:"name_1"`
	Name2     string    `json:"name_2"`
	Address   string    `json:"address"`
	Postcode  string    `json:"postcode"`
	City      string    `json:"city"`
	Mail      string    `json:"mail"`
	Phone     string    `json:"phone_fixed"`
	Mobile    string    `json:"phone_mobile"`
	UpdatedAt Timestamp `json:"updated_at"`
}

func parseListContactsResponse(b []byte) ([]Contact, error) {
	var contacts []Contact

	err := json.Unmarshal(b, &contacts)
	if err != nil {
		return []Contact{}, err
	}

	return contacts, nil
}

func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.Name, c.Name2)
}
