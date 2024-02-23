package domain

import (
	"testing"
)

func TestNewOrderRepositoryDb(t *testing.T) {
	repo := NewOrderRepositoryDb()
	if repo.client == nil {
		t.Error("Expected a non-nil client in the repository")
	}

}
