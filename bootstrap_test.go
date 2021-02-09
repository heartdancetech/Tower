package tower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBootStrap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBootStrap(nil)
			assert.NotNil(t, b)
		})
	}
}

func Test_bootStrap_SetOnConnSatrt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{}
			c.check()
			b := &bootStrap{
				Config:      &c,
				ConnMgr:     nil,
				OnConnStart: nil,
				OnConnClose: nil,
			}
			b.SetOnConnStart(func(conn Connectioner) {})
			assert.NotNil(t, b.OnConnStart)
		})
	}
}

func Test_bootStrap_SetOnConnClose(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{}
			c.check()
			b := &bootStrap{
				Config:      &c,
				ConnMgr:     nil,
				OnConnStart: nil,
				OnConnClose: nil,
			}
			b.SetOnConnClose(func(conn Connectioner) {})
			assert.NotNil(t, b.OnConnClose)
		})
	}
}
