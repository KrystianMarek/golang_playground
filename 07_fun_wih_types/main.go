package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"
)

func print(writer *tabwriter.Writer, stuff any) {
	fmt.Fprintln(writer, fmt.Sprintf("%v:\t%v", reflect.TypeOf(stuff), stuff))
}

func lineBreak(writer *tabwriter.Writer) {
	line := ""
	for i := 0; i < 10; i++ {
		line += "_"
	}
	fmt.Fprintln(writer, fmt.Sprintf("%s\t%s", line, line))
}

type SimpleStruct struct {
	Username string
	Password string
	Id       int64
}

func playgorund(writer *tabwriter.Writer) {
	yo := "Yo!"
	print(writer, yo)
	print(writer, &yo)
	lineBreak(writer)

	one := 1
	print(writer, one)
	lineBreak(writer)

	oneThird := 0.3333333333333333333333333
	print(writer, oneThird)
	print(writer, oneThird*3)
	lineBreak(writer)

	mySlice := []int{1, 2, 3, 4}
	print(writer, mySlice)
	print(writer, mySlice[0])
	lineBreak(writer)

	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	print(writer, timeZone)
	lineBreak(writer)

	printFunction := print
	print(writer, printFunction)
	print(writer, &printFunction)
	lineBreak(writer)

	print(writer, SimpleStruct{})
	print(writer, new(SimpleStruct))
	simpleStruct := SimpleStruct{Id: 12345, Username: "John", Password: "Doe"}
	print(writer, simpleStruct)
	print(writer, &simpleStruct)
	ssPtr := &simpleStruct
	print(writer, ssPtr)
	print(writer, *ssPtr)
	print(writer, &ssPtr)
	lineBreak(writer)

	myError := errors.New("oops")
	print(writer, myError)
	lineBreak(writer)

}

func main() {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer writer.Flush()
	playgorund(writer)
}
