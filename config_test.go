package tower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_check(t *testing.T) {

	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				Name:             "",
				IP:               "",
				Port:             0,
				MaxPacketSize:    0,
				MaxConn:          0,
				WorkerPoolSize:   0,
				MaxWorkerTaskLen: 0,
				MaxMsgChanLen:    0,
				Logging:          nil,
			}
			c.check()
			//assert.NotEqual(t, "", c.Name)
			assert.NotEqual(t, c.IP, "")
			assert.NotEqual(t, c.Port, 0)
			//assert.NotEqual(t, c.MaxConn, 0)
			//assert.NotEqual(t, c.WorkerPoolSize, 0)
			//assert.NotEqual(t, c.MaxPacketSize, 0)
			//assert.NotEqual(t, c.MaxWorkerTaskLen, 0)
			//assert.NotEqual(t, c.MaxMsgChanLen, 0)
			assert.NotNil(t, c.Logging)
		})
	}
}
