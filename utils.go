package go_aoahid

import (
	"encoding/binary"
	"github.com/google/gousb"
)

func getProtocol(dev *gousb.Device) (uint16, error) {
	if dev == nil {
		return 0, ErrorNoAccessoryDevice
	}
	data := make([]byte, 2)
	_, err := dev.Control(RTypeIn, AccessoryGetProtocol, 0, 0, data)
	if err != nil {
		return 0, ErrorFailedToGetProtocol
	}
	protocol := binary.BigEndian.Uint16(data)
	return protocol, nil
}
