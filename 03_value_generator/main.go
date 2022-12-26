package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
)

func getSliceOfRandomInt(capacity int, maxIncrement int) []int {
	myMap := make(map[int]int)

	for index := 0; index < capacity; index++ {
		value := index + rand.Intn(maxIncrement)
		myMap[value] = value
	}

	result := make([]int, 0, len(myMap))
	for _, value := range myMap {
		result = append(result, value)
	}

	sort.Ints(result)
	return result
}

func writeSliceOfInt(capacity int) {
	data, _ := json.Marshal(getSliceOfRandomInt(capacity, 42))

	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)

	err := os.WriteFile(filepath.Join(exPath, fmt.Sprintf("../data/sliceOfInt_%d.json", capacity)), data, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	writeSliceOfInt(10)
	writeSliceOfInt(10000)
}
