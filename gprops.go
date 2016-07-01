// Written 2015 by Marcin 'Zbroju' Zbroinski.
// Use of this source code is governed by GNU General Public License
// that can be found in the LICENCE file.

/*
Package gprops implements simple properties object, similar
to the one known from java.

It can be used to store and load simple configuration data in a form
of key = value pair.

All lines beginning with '#' are omitted - assuming they are comments.
*/
package gprops

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

const (
	commentPrefix    = "#"
	settingSeparator = "="
)

// Properties type is an object containing properties.
type Properties struct {
	propertiesMap map[string]string
}

// New method creates a new empty Props object.
func New() *Properties {
	return &Properties{propertiesMap: make(map[string]string)}
}

// Set method sets new value for given key. If the key doesn't exist, it will be created.
func (props *Properties) Set(key, value string) {
	props.propertiesMap[key] = value
}

// Get method returns property value for the key
func (props *Properties) Get(key string) string {
	return props.propertiesMap[key]
}

// GetOrDefault method returns property value for the key if it exist, or default value if otherwise.
func (props *Properties) GetOrDefault(key, defaultValue string) string {
	if props.Contains(key) {
		return props.propertiesMap[key]
	} else {
		return defaultValue
	}
}

// Contains method returns true if given key exists.
func (props *Properties) Contains(key string) bool {
	_, exists := props.propertiesMap[key]
	return exists
}

// Delete method removes existing property.
func (props *Properties) Delete(key string) {
	delete(props.propertiesMap, key)
}

// Load loads properties from a reader (e.g. config file) to the properties and return error in case of problems.
// Lines with '#' at the beginning are skipped.
func (props *Properties) Load(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, commentPrefix) && len(line) > 0 {
			keyValuePair := strings.SplitN(line, settingSeparator, 2)
			if len(keyValuePair) == 2 {
				props.Set(strings.TrimSpace(keyValuePair[0]), strings.TrimSpace(keyValuePair[1]))
			} else {
				return errors.New("incorrect syntax in config file.")
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
func (props *Properties) Store(w io.Writer, comment string) error {
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	if comment != "" {
		bw.WriteString(commentPrefix + " " + comment + "\n")
	}
	for key, value := range props.propertiesMap {
		_, err := bw.WriteString(key + settingSeparator + value + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
