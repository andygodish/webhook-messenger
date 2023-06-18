package main

import (
	"net/http"
	"testing"
)

type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestSendMessage(t *testing.T) {
	doCalled := false
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			doCalled = true
			return nil, nil
		},
	}	

	messenger := TeamsMessenger{
		Client: mockClient,
	}

	err := messenger.SendMessage("http://example.com", TeamsMessageCard{Text: "hello"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !doCalled {
		t.Error("Expected Do to be called, but it wasn't")
	}
}
