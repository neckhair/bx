package bexio_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/neckhair/bx/bexio"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

type mockTokenSource struct{}

func (ts *mockTokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{AccessToken: "fake-token"}
	return token, nil
}

func newTestClient(url string) *bexio.Client {
	tokenSource := &mockTokenSource{}
	client := bexio.NewClient(context.Background(), tokenSource)
	client.BaseUrl = url
	return client
}

func readTestJsonFromFile(t *testing.T, filename string) string {
	path := filepath.Join("testdata", filename)
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("could not open test data file: %v", err)
		return ""
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
		return ""
	}
	return string(data)
}

func TestContactFullName(t *testing.T) {
	contact := bexio.Contact{Name: "Meier", Name2: "Herbert"}
	assert.Equal(t, contact.FullName(), "Meier Herbert")
}

func TestListContacts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		testdata := readTestJsonFromFile(t, "list_contacts.json")
		fmt.Fprint(w, string(testdata))
	}))
	defer ts.Close()

	client := newTestClient(ts.URL)

	contacts, err := bexio.ListContacts(client, 100)

	assert.NoError(t, err)
	assert.Equal(t, "Meier", contacts[1].Name)
	assert.Equal(t, "Sepp", contacts[1].Name2)
	assert.Equal(t, "2020-01-09 12:18:15 +0000 UTC", contacts[1].UpdatedAt.String())
}

func TestListContactsNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "")
	}))
	defer ts.Close()

	client := newTestClient(ts.URL)

	contacts, err := bexio.ListContacts(client, 100)

	assert.ErrorAs(t, err, &bexio.NotFoundError)
	assert.Empty(t, contacts)
}

func TestListContactsUnauthorized(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "not authorized")
	}))
	defer ts.Close()

	client := newTestClient(ts.URL)

	contacts, err := bexio.ListContacts(client, 100)

	assert.ErrorAs(t, err, &bexio.UnauthorizedError)
	assert.Empty(t, contacts)
}

func TestListContactsServerError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "internal server error")
	}))
	defer ts.Close()

	client := newTestClient(ts.URL)

	contacts, err := bexio.ListContacts(client, 100)

	assert.Error(t, err)
	assert.Empty(t, contacts)
}
