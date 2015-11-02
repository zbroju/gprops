package gprops_test

import (
	"github.com/zbroju/gprops"
	"testing"
)

func TestNew(t *testing.T) {
	properties := gprops.New()
	properties.Add("key1", "value1")
}

func TestAdd(t *testing.T) {
	properties := gprops.New()
	properties.Add("key1", "value1")
}

func TestGet(t *testing.T) {
	properties := gprops.New()
	properties.Add("key1", "value1")
	if properties.Get("key1") != "value1" {
		t.Errorf("Method 'Get' returns a different value than expected.\n")
	}
}
