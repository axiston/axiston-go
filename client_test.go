package axiston_test

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	if _, err := NewClient("testing",
		WithUserAgent("Axiston/1.0"),
	); err != nil {
		t.Error(err)
	}
}

func TestClient_Health(t *testing.T) {
	client, _ := NewClient("testing")
	if err := client.Health(); err != nil {
		t.Error(err)
	}
}
