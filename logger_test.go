package tower

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var _ Logger = new(logging)
		})
	}
}

func Test_logging_LogMode(t *testing.T) {
	tests := []struct {
		name string
		mode LogLevel
	}{
		{"test_debug", Debug},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var testLogging = NewLogger(log.New(os.Stdout, "[Tower] ", log.Ldate|log.Ltime|log.LUTC), tt.mode)
			testLogging.LogMode(Debug)
			assert.NotNil(t, testLogging)
		})
	}
}
