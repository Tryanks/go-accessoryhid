package accessory

import (
	"testing"
)

func TestGetProtocol(t *testing.T) {
	devices, err := GetDevices(1)
	if err != nil {
		t.Error(err)
	}
	for _, v := range devices {
		t.Log(v.Protocol)
		_ = v.Close()
	}
	device, err := GetDevice()
	if err != nil {
		t.Error(err)
	}
	_ = device.Close()
	device, err = GetDeviceWithSerial("")
	if err != nil {
		t.Error(err)
	}
	_ = device.Close()
}
