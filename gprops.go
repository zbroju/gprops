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

//TODO: Load(reader reader)
//TODO: Store(writer writer, comments string)
