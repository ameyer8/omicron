package model

import (
	"fmt"
	"testing"
)

func TestSetPortName(t *testing.T) {
	setPortName(1, "testName")
	setPortName(2, "testName2")
	setPortName(1, "this is an led")

	name, err := getPortName(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Name of Port 1 is %s", name)

}
