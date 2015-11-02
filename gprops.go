// Copyright 2015 Marcin 'Zbroju' Zbroinski. All rights reserved.
// Use of this source code is governed by GNU General Public License
// that can be found in the LICENCE file.

/*
Package gprops implements simple properties object, similar
to the one know from java.
*/
package gprops

// Props type is an object containing properties.
type Props struct {
	propsMap map[string]string
}

// New method creates a new empty Props object.
func New() *Props {
	return &Props{propsMap: make(map[string]string)}
}

// Add new property
func (props *Props) Add(key, value string) {
	props.propsMap[key] = value
}

// Get property value for the key
func (props *Props) Get(key string) string {
	return props.propsMap[key]
}
