package bmc

import (
	"context"

	"github.com/ihciah/bmc/pkg/ipmi"
)

// SessionCommands contains high-level wrappers for sending commands within an
// established session. These commands are common to all versions of IPMI.
type SessionCommands interface {

	// All session-less commands can also be sent inside a session; indeed it is
	// convention for Get Channel Authentication Capabilities to be used as a
	// keepalive.
	SessionlessCommands

	// GetDeviceID send a Get Device ID command to the BMC. This is specified in
	// 17.1 and 20.1 of IPMI v1.5 and 2.0 respectively.
	GetDeviceID(context.Context) (*ipmi.GetDeviceIDRsp, error)

	// GetChassisStatus sends a Get Chassis Status command to the BMC. This is
	// specified in 22.2 and 28.2 of IPMI v1.5 and 2.0 respectively.
	GetChassisStatus(context.Context) (*ipmi.GetChassisStatusRsp, error)

	// ChassisControl provides power up, power down and reset control. It is
	// specified in 22.3 and 28.3 of IPMI v1.5 and 2.0 respectively.
	ChassisControl(context.Context, ipmi.ChassisControl) error

	// GetSDRRepositoryInfo obtains information about the BMC's Sensor Data
	// Record Repository. It is specified in 27.9 and 33.9 of IPMI v1.5 and 2.0
	// respectively.
	GetSDRRepositoryInfo(context.Context) (*ipmi.GetSDRRepositoryInfoRsp, error)

	// GetSensorReading retrieves the current value of a sensor, identified by
	// its number. It is specified in 29.14 and 35.14 of IPMI v1.5 and 2.0
	// respectively. Note, the raw value is in one of three formats, and is
	// converted into a "real" reading via one or more formulae - interpreting
	// it requires the SDR.
	GetSensorReading(context.Context, uint8) (*ipmi.GetSensorReadingRsp, error)

	// closeSession sends a Close Session command to the BMC. It is unexported
	// as calling it randomly would leave the session in an invalid state. Call
	// Close() on the session itself to invoke this.
	closeSession(context.Context) error
}
