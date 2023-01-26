package main

import (
	"fmt"
	"fun_wih_types/channel"
	"fun_wih_types/pipe"
	"fun_wih_types/read_write"
	"fun_wih_types/stdout"
	"reflect"
)

type Demonstrable interface {
	Demo()
}

// https://stackoverflow.com/questions/35790935/using-reflection-in-go-to-get-the-name-of-a-struct
func getType(it interface{}) string {
	t := reflect.TypeOf(it)

	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func main() {
	demos := []Demonstrable{
		channel.SimpleChannel{},
		pipe.PipeExample{},
		read_write.ReadWriteExample{},
		stdout.StdoutPipeExample{},
	}

	for _, demo := range demos {
		componentName := getType(demo)
		fmt.Printf("#### BEGIN:%s ####\n", componentName)
		demo.Demo()
		fmt.Printf("####  END:%s  ####\n", componentName)
	}
}
