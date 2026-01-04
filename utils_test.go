package accessory

import (
	"testing"

	"github.com/google/gousb"
)

func TestGetProtocol(t *testing.T) {
	ctx := gousb.NewContext()
	defer ctx.Close()
	devices, err := GetDevices(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	for _, v := range devices {
		t.Log(v.Protocol)
		_ = v.Close()
	}
	device, err := GetDevice(ctx)
	if err != nil {
		t.Error(err)
	} else {
		_ = device.Close()
	}
	device, err = GetDeviceWithSerial(ctx, "")
	if err != nil {
		t.Error(err)
	} else {
		_ = device.Close()
	}
}
