package tower

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var (
	testLogger = NewLogger(log.New(os.Stdout, "[Tower] ", log.Ldate|log.Ltime|log.LUTC), Debug)
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var _ logger = new(logging)
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
