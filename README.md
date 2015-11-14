# gprops

## Description

Package `gprops` implements simple properties object for applications written in `golang`, similar to the one known from java.

It can be used to store and load simple configuration data in a form of pairs:
```
key = value
```
Bear in mind that both `key` and `value` are strings. All lines beginning with `#` are omitted - assuming they are comments.

## Installation

This package can be installed with the go get command:
``` go
go get github.com/zbroju/gprops
```
## Documentation

### Example of configuration file
```
# Example of configuration file made by gprops package

# Settings are written in a simple KEY = VALUE pairs:
DATA_FILE = /home/.examplerc
VERBOSE = 1

# All lines starting with '#' are skipped, assuming they are just comments
```
### Usage
Remember to include gprops package in your application:
``` go
import "github.com/zbroju/gprops"
```
#### Load properties from file
``` go
// Open config file
file, err := os.Open(".configFile")
if err != nil {
    fmt.Println(err)
}
defer file.Close()

// Try to load properties and read the data
properties := gprops.NewProps()
errLoad := properties.Load(file)
if errLoad != nil {
    fmt.Println(errLoad.Error())
}

// Assuming the config file looks like the example above
file := properties.Get("DATA_FILE") // variable 'file' contains "/home/.examplerc" value.
verboseFlag := properties.Get("VERBOSE") // variable 'verboseFlag' contains "1" value.
```
#### Save properties to a file
``` go
// Prepare properties object
properties := gprops.NewProps()
properties.Set("DATA_FILE", "/home/.examplerc")
properties.Set("VERBOSE", "1")

// Create new file with properties
f, err := os.Create(".configFile")
if err != nil {
    fmt.Println(err.Error())
}

// Store properties in the file
propsToStore.Store(f, "Example of configuration file made by gprops package")
f.Close()
```
## License

GNU General Public License.

## Author
Marcin 'Zbroju' Zbroiński