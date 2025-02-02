package db

import (
	"testing"
)

func TestGetDatabaseConnection(t *testing.T) {
	// Get env variables
	instance, err := GetDatabaseConnection()
	if err != nil {
		t.Fatalf("ErrorGettingConnection: %v", err)
	}

	if instance == nil {
		t.Fatal("EmptyDatabaseInstance")
	}

	pingErr := instance.Ping()
	if pingErr != nil {
		t.Fatalf("Ping Error \n%v", pingErr)
	}
}
