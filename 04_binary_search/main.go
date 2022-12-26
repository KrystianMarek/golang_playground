package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func readData(path string) []int {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)

	fileContent, err := os.ReadFile(filepath.Join(exPath, path))
	if err != nil {
		panic(err)
	}

	var sliceOfRandomInt []int
	err = json.Unmarshal(fileContent, &sliceOfRandomInt)
	if err != nil {
		panic(err)
	}

	return sliceOfRandomInt
}

func binarySearch(data []int, target int, start int, stop int) (int, int) {
	if stop < start {
		fmt.Printf("\tValue: %d not found, closest value: %d|%d\n", target, data[start], data[stop])
		return -1, -1
	}
	middle := int((start + stop) / 2)

	//fmt.Printf("\ttarget: v%d, middle: k%d|v%d, start: k%d|v%d, stop: k%d|v%d\n", target, middle, data[middle], start, data[start], stop, data[stop])

	if data[middle] > target {
		return binarySearch(data, target, start, middle)
	} else if data[middle] < target {
		return binarySearch(data, target, middle+1, stop)
	} else {
		return middle, data[middle]
	}
}

func find(data []int, target int) {
	index, val := binarySearch(data, target, 0, len(data)-1)
	fmt.Printf("Index of %d(%d) is %d\n", target, val, index)
}

func main() {
	data := readData("../resources/sliceOfInt_10000.json")
	fmt.Printf("min: %d, max: %d, len: %d\n", data[0], data[len(data)-1], len(data))

	find(data, 1234)

	find(data, 20)

	find(data, 7435)
}
