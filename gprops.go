// Copyright 2015 Marcin 'Zbroju' Zbroinski. All rights reserved.
// Use of this source code is governed by GNU General Public License
// that can be found in the LICENCE file.

/*
Package gprops implements simple properties object, similar
to the one known from java.

It can be used to store and load simple configuration data in a form
of key = value pair.
*/
package gprops

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// Props type is an object containing properties.
type Props struct {
	propsMap map[string]string
}

// New method creates a new empty Props object.
func NewProps() *Props {
	return &Props{propsMap: make(map[string]string)}
}

// Set method sets new value for given key. If the key doesn't exist, it will be created.
func (props *Props) Set(key, value string) {
	props.propsMap[key] = value
}

// Get method returns property value for the key
func (props *Props) Get(key string) string {
	return props.propsMap[key]
}

// ContainsKey method returns true if given key exists.
func (props *Props) ContainsKey(key string) bool {
	_, exists := props.propsMap[key]
	return exists
}

// Delete method removes existing property.
func (props *Props) Delete(key string) {
	delete(props.propsMap, key)
}

// Load loads properties from a reader (e.g. config file) to the properties and return error in case of problems.
// Lines with '#' at the beginning are skipped.
func (props *Props) Load(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			keyValuePair := strings.Split(line, "=")
			if len(keyValuePair) == 2 {
				props.Set(strings.TrimSpace(keyValuePair[0]), strings.TrimSpace(keyValuePair[1]))
			} else {
				return errors.New("gprops: incorrect syntax in input data.")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	} else {
		return nil
	}
}

// Store stores properties using a given writer (e.g. config file) and return in case of problems.
// You can add comment, which will be stored as first line beginning with '#'.
func (props *Props) Store(w *bufio.Writer, comment string) error {
	if comment != "" {
		w.WriteString("# " + comment + "\n")
	}
	for key, value := range props.propsMap {
		_, err := w.WriteString(key + "=" + value + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
