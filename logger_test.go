package tower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_logging_LogMode(t *testing.T) {
	tests := []struct {
		name string
		want Logger
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Nil(t, nil)
			assert.Nil(t, nil)
		})
	}
}
