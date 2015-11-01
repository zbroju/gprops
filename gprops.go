// Copyright 2015 Marcin 'Zbroju' Zbroinski. All rights reserved.
// Use of this source code is governed by GNU General Public License
// that can be found in the LICENCE file.

/*
Package gprops implements simple properties object, similar
to java one.
*/
package gprops

// Props type is an object containing properties.
type Props map[string]string

// New method creates a new empty Props object.
func New() Props {
	return make(Props)
}
