package bexio_test

import (
	"net/http"
	"testing"

	"github.com/neckhair/bx/bexio"
	"github.com/neckhair/bx/bexio/mocks"
	"github.com/stretchr/testify/assert"
)

func TestContactFullName(t *testing.T) {
	contact := bexio.Contact{Name: "Meier", Name2: "Herbert"}
	assert.Equal(t, contact.FullName(), "Meier Herbert")
}

func TestListContacts(t *testing.T) {
	ts := mocks.NewTestServer(http.StatusOK, mocks.ContactList())
	defer ts.Close()

	client := newTestClient(ts.URL)

	contacts, err := client.ListContacts(100)

	assert.NoError(t, err)
	assert.Equal(t, "Meier", contacts[1].Name)
	assert.Equal(t, "Sepp", contacts[1].Name2)
	assert.Equal(t, "2020-01-09 12:18:15 +0000 UTC", contacts[1].UpdatedAt.String())
}

func TestListContactsNotFound(t *testing.T) {
	ts := mocks.NewTestServer(http.StatusNotFound, "")
	defer ts.Close()

	client := newTestClient(ts.URL)

	contacts, err := client.ListContacts(100)

	assert.ErrorAs(t, err, &bexio.NotFoundError)
	assert.Empty(t, contacts)
}

func TestListContactsUnauthorized(t *testing.T) {
	ts := mocks.NewTestServer(http.StatusUnauthorized, "not authorized")
	defer ts.Close()

	client := newTestClient(ts.URL)

	contacts, err := client.ListContacts(100)

	assert.ErrorAs(t, err, &bexio.UnauthorizedError)
	assert.Empty(t, contacts)
}

func TestListContactsServerError(t *testing.T) {
	ts := mocks.NewTestServer(http.StatusInternalServerError, "internal server error")
	defer ts.Close()

	client := newTestClient(ts.URL)

	contacts, err := client.ListContacts(100)

	assert.Error(t, err)
	assert.Empty(t, contacts)
}
