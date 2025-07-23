package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var (
	wg sync.WaitGroup 
	mu sync.Mutex
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
	max := math.MinInt

	for _, val := range data {
		//if max < val { max = val} ??
		max = int(math.Max(float64(val), float64(max)))
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь

	return maximum(data)
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

	sizeSlice := int(math.Ceil(float64(len(data)) / float64(CHUNKS)))
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
			
		currentIdx := i * sizeSlice
		endIdx := currentIdx + sizeSlice
		if endIdx > len(data) {
			endIdx = len(data)
		}
		chunk := data[currentIdx : endIdx]


		go func(){
			defer wg.Done()
			maxLocal := maxChunks(chunk)
			mu.Lock()
			if(max < maxLocal){
				max = maxLocal
			}
			mu.Unlock()
			
		}()
	}

	wg.Wait()
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
