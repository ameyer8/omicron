package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ameyer8/omicron/internal/data"
)

var portNameFile string = "/opt/omicron/portName.json"

func getPortName(portID int) (string, error) {

	file, err := os.Open(portNameFile)
	defer file.Close()
	if err != nil {
		log.Printf("[getPortName] Unable to open %s\n", portNameFile)
		return "", err
	}
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("[getPortName] Unable to read %s\n", portNameFile)
		return "", err
	}
	var portNames []data.PortName
	json.Unmarshal(fileData, &portNames)

	var portName data.PortName
	for _, pn := range portNames {
		if pn.ID == portID {
			portName = pn
			break
		}
	}

	return portName.Name, nil

}

func setPortName(portID int, name string) error {
	var err error

	file, err := os.Open(portNameFile)
	defer file.Close()
	if err != nil {
		log.Printf("[getPortName] Unable to open %s\n", portNameFile)
		return err
	}
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("[getPortName] Unable to read %s\n", portNameFile)
		return err
	}
	var portNames []data.PortName
	json.Unmarshal(fileData, &portNames)

	portFound := false
	for idx, pn := range portNames {
		if pn.ID == portID {
			portNames[idx].Name = name
			portFound = true
			break
		}
	}
	if !portFound {
		portNames = append(portNames, data.PortName{ID: portID, Name: name})
	}
	tmpFile, err := ioutil.TempFile(os.TempDir(), "omicron")
	if err != nil {
		log.Printf("[getPortName] Unable to create file at %s\n", os.TempDir()+tmpFile.Name())
		return err
	}

	fileData, err = json.Marshal(portNames)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(fileData))
	os.Truncate(portNameFile, 0)
	err = ioutil.WriteFile(portNameFile, fileData, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	return err
}
