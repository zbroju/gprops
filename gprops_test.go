package gprops_test

import (
	"github.com/zbroju/gprops"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	properties := gprops.NewProps()
	properties.Set("key1", "value1")
}

func TestSet(t *testing.T) {
	properties := gprops.NewProps()
	properties.Set("key1", "value1")
}

func TestGet(t *testing.T) {
	properties := gprops.NewProps()
	properties.Set("key1", "value1")
	if properties.Get("key1") != "value1" {
		t.Errorf("Returned value is not as expected.\n")
	}
}

func TestContainsKey(t *testing.T) {
	properties := gprops.NewProps()
	properties.Set("key", "")
	if properties.ContainsKey("key") == false {
		t.Errorf("Returns false for existing key.\n")
	}
	if properties.ContainsKey("anotherkey") == true {
		t.Errorf("Returns true for not existing key.\n")
	}
}

func TestDelete(t *testing.T) {
	properties := gprops.NewProps()
	properties.Set("key1", "value1")
	properties.Delete("key1")
	if properties.ContainsKey("key1") == true {
		t.Errorf("The key has not been deleted.\n")
	}
}

func TestLoad(t *testing.T) {
	// Create temporary data & file
	tempFile := "tempRCData"

	keyAndValues := []string{"# Example config file", "FILE", "/home/user/myfile", "VERBOSE", "true"}
	byteStream := []byte(keyAndValues[0] + "\n" + keyAndValues[1] + " = " + keyAndValues[2] + "\n" + keyAndValues[3] + " = " + keyAndValues[4] + "\n")

	err := ioutil.WriteFile(tempFile, byteStream, 0644)
	if err != nil {
		t.Errorf("Problem with creating temporary file.")
	}
	defer os.Remove(tempFile)

	// Open temporary file
	file, err := os.Open(tempFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Try to load properties and read the data
	properties := gprops.NewProps()
	errLoad := properties.Load(file)
	if errLoad != nil {
		t.Errorf(errLoad.Error())
	}
	if properties.Get(keyAndValues[1]) != keyAndValues[2] || properties.Get(keyAndValues[3]) != keyAndValues[4] {
		t.Errorf("Data loaded from file are not as expected.\n")
	}
}
