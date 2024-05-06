package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name    string
		dbName  string
		wantErr assert.ErrorAssertionFunc
	}{
		// Test cases
		{"Valid In-Memory Database", ":memory:", assert.NoError},
		{"Invalid File Database", "/invalid/file/path", assert.Error},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Connect(tt.dbName)
			tt.wantErr(t, err)
		})
	}
}
