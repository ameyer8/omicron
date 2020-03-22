package model

import (
	"errors"

	onion "github.com/ameyer8/omicron/external/omega2gpio"
)

//GetGPIODirection return direction of the GPIO port
func GetGPIODirection(portID int) string {

	var retStr string
	switch onion.GetDirection(portID) {
	case 0:
		retStr = "in"
	case 1:
		retStr = "out"
	default:
	}
	return retStr
}

//SetGPIODirection return direction of the GPIO port
func SetGPIODirection(portID int, dir string) error {

	var err error
	switch dir {
	case "in":
		onion.SetDirection(portID, 0)
	case "out":
		onion.SetDirection(portID, 1)
	default:
		err = errors.New("Could not set direction. Direction not recognized")
	}
	return err
}

//GetGPIOValue reads the voltage level of the port
func GetGPIOValue(portID int) int {

	level := onion.Read(portID)
	return int(level)
}

//SetGPIOValue sets level of the GPIO port
func SetGPIOValue(portID int, value int) error {

	var err error
	if GetGPIODirection(portID) != "out" {
		err = errors.New("Port not set to output mode. Cannot set")
		return err
	}
	switch value {
	case 0, 1:
		// Check current value
		if v := GetGPIOValue(portID); v != value {
			onion.Write(portID, uint8(value))
		}
	default:
		err = errors.New("Could not set direction. Direction not recognized")
	}
	return err
}
func ToggleGPIOValue(portID int) error {

	var err error
	if v := GetGPIOValue(portID); v == 0 {
		err = SetGPIOValue(portID, 1)
	} else {
		err = SetGPIOValue(portID, 0)
	}

	return err
}
