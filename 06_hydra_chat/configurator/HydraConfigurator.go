package configurator

import (
	"errors"
	"fmt"
	"reflect"
)

type Configuration struct {
	Name       string `name:"name"`
	RemoteAddr string `name:"remoteip"`
	TCP        bool   `name:"tcp"`
}

var wrongTypeError error = errors.New("type must be a pointer to a struct")

func (config Configuration) GetConfiguration(filename string) (err error) {
	//check if this is type pointer

	mysRValue := reflect.ValueOf(&config)
	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
		return wrongTypeError
	}
	//get and confirm the struct value
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return wrongTypeError
	}

	err = MarshalCustomConfig(mysRValue, filename)

	fmt.Println(config.RemoteAddr)

	return err
}

func NewConfiguration() (config Configuration) {
	return Configuration{}
}
