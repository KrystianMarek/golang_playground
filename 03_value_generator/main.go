package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
)

func getRandomSliceOfInt(capacity int, maxIncrement int) []int {
	var result []int

	for index := 0; index < capacity; index++ {
		result = append(result, index+rand.Intn(maxIncrement))
	}

	return result
}

func writeSliceOfInt() {
	data, _ := json.Marshal(getRandomSliceOfInt(10000, 42))

	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)

	err := os.WriteFile(filepath.Join(exPath, "../data/sliceOfInt.json"), data, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	writeSliceOfInt()
}
