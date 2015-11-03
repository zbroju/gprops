package gprops_test

import (
	"github.com/zbroju/gprops"
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
		t.Errorf("Method 'Get' returns a different value than expected.\n")
	}
}

func TestContainsKey(t *testing.T) {
	properties := gprops.NewProps()
	properties.Set("key", "")
	if properties.ContainsKey("key") == false {
		t.Errorf("Method 'ContainsKey' doesn't return true for existing key.\n")
	}
	if properties.ContainsKey("anotherkey") == true {
		t.Errorf("Method 'ContainsKey' returns true for non existing key.\n")
	}
}

func TestDelete(t *testing.T) {
	properties := gprops.NewProps()
	properties.Set("key1", "value1")
	properties.Delete("key1")
	if properties.ContainsKey("key1") == true {
		t.Errorf("Method 'Delete' does not delete existing key.\n")
	}
}
