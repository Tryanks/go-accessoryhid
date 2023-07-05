package go_aoahid

import (
	"github.com/google/gousb"
	"testing"
)

func TestGetProtocol(t *testing.T) {
	devices, err := gousb.NewContext().OpenDevices(func(desc *gousb.DeviceDesc) bool {
		return true
	})
	if err != nil {
		t.Error(err)
	}
	for _, v := range devices {
		p, err := getProtocol(v)
		if err != nil {
			t.Error(err)
		}
		t.Log(p)
	}
}
