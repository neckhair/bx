package bexio

import "fmt"

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

func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.Name, c.Name2)
}
