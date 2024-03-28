package accessory

import (
	"testing"
)

func TestGetProtocol(t *testing.T) {
	devices, err := getDevices(1, true)
	if err != nil {
		t.Error(err)
	}
	for _, v := range devices {
		t.Log(v.Protocol)
		_ = v.Close()
	}

	devices, err = GetDevices(1)
	if err != nil {
		t.Error(err)
	}
	for _, v := range devices {
		t.Log(v.Protocol)
		_ = v.Close()
	}
}
