package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Set map[int]int

func (s Set) getValues() []int {
	var result []int
	for _, value := range s {
		result = append(result, value)
	}

	sort.Ints(result)
	return result
}

func getSliceOfRandomInt(capacity int, maxIncrement int) []int {
	set := make(Set)
	rand.Seed(time.Now().UnixNano())

	for index := 0; index < capacity; index++ {
		value := index + rand.Intn(maxIncrement)
		set[value] = value
	}

	return set.getValues()
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
