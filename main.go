package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	var slice []int

	if size == 0 {
		return slice;
	}

	for i := 0; i < size; i++ {
		slice = append(slice, rand.Int());
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if(len(data) == 0){
		return 0
	}
	// ваш код здесь
	max := slices.Max(data)

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь

	var wg sync.WaitGroup 

	//Если массив меньше 100_000_000

	sizeChunk := CHUNKS
	if CHUNKS > len(data){
		sizeChunk = len(data)
	}

	wg.Add(sizeChunk)
	resultsMax := make([]int, sizeChunk)

	sizeSlice := len(data) / sizeChunk

	for i := 0; i < sizeChunk; i++ {
		currentIdx := i * sizeSlice
		endIdx := currentIdx + sizeSlice
		if endIdx > len(data) {
			endIdx = len(data)
		}
		chunk := data[currentIdx : endIdx]

		go func(chunk []int, i int){
			defer wg.Done()
			maxLocal := maximum(chunk)
			resultsMax[i] = maxLocal
		}(chunk, i)
	}

	wg.Wait()

	return slices.Max(resultsMax)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	var max int
	data := generateRandomElements(SIZE)

	fmt.Printf("Ищем максимальное значение в один поток\n")
	// ваш код здесь
	start := time.Now()
	max = maximum(data)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
