package configurator

import (
	"errors"
	"reflect"
)

type Configuration struct {
	Name       string `name:"name"`
	RemoteAddr string `name:"remoteip"`
	TCP        bool   `name:"tcp"`
}

var wrongTypeError error = errors.New("type must be a pointer to a struct")

func (config *Configuration) GetConfiguration(filename string) error {
	//check if this is type pointer

	mysRValue := reflect.ValueOf(config)
	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
		return wrongTypeError
	}
	//get and confirm the struct value
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return wrongTypeError
	}

	err := MarshalCustomConfig(mysRValue, filename)

	return err
}
