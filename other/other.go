/*
Package other is an example of a custom driver
*/
package other

import (
	"github.com/urionz/config"
	"github.com/urionz/config/ini"
)

const driverName = "other"

var (
	// Driver is the exported symbol
	Driver = &otherDriver{}
	// Encoder is the encoder for this driver
	Encoder = ini.Encoder
	// Decoder is the decoder for this driver
	Decoder = ini.Decoder
)

type otherDriver struct{}

// Name get name
func (d *otherDriver) Name() string {
	return driverName
}

// GetDecoder for other (same than ini)
func (d *otherDriver) GetDecoder() config.Decoder {
	return Decoder
}

// GetEncoder for other (same than ini)
func (d *otherDriver) GetEncoder() config.Encoder {
	return Encoder
}
