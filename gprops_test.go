package gprops_test

import (
	"github.com/zbroju/gprops"
	"io/ioutil"
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

	comment := "# Example config file"
	key1 := "FILE"
	val1 := "/home/user/myfile"
	key2 := "VERBOSE"
	val2 := "true"
	byteStream := []byte(comment + "\n" + key1 + " = " + val1 + "\n" + key2 + " = " + val2 + "\n")

	err := ioutil.WriteFile(tempFile, byteStream, 0644)
	if err != nil {
		t.Errorf("Problem with creating temporary file.")
	}
	defer os.Remove(tempFile)

	// Open temporary file
	file, err := os.Open(tempFile)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer file.Close()

	// Try to load properties and read the data
	properties := gprops.NewProps()
	errLoad := properties.Load(file)
	if errLoad != nil {
		t.Errorf(errLoad.Error())
	}
	if properties.Get(key1) != val1 || properties.Get(key2) != val2 {
		t.Errorf("Data loaded from file are not as expected.\n")
	}
}

func TestStore(t *testing.T) {
	// Prepare the properties to store
	propsToStore := gprops.NewProps()
	key1 := "key1"
	val1 := "val1"
	key2 := "key2"
	val2 := "val2"
	propsToStore.Set(key1, val1)
	propsToStore.Set(key2, val2)

	// Store properties in file
	tempFile := "temporaryPropsFile"
	f, err := os.Create(tempFile)
	if err != nil {
		t.Errorf(err.Error())
	}
	propsToStore.Store(f, "Test properties")
	f.Close()

	// Load the properties from the file
	propsLoaded := gprops.NewProps()
	f2, err := os.Open(tempFile)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer f2.Close()
	errLoad := propsLoaded.Load(f2)
	if errLoad != nil {
		t.Errorf(errLoad.Error())
	}

	// Compare the properties
	if propsLoaded.Get(key1) != propsToStore.Get(key1) || propsLoaded.Get(key2) != propsToStore.Get(key2) {
		t.Errorf("Data stored in a file are not as expected.\n")
	}
	os.Remove(tempFile)
}
