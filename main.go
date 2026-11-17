package main

import (
	"fmt"
	"math/rand"
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
	if size <= 0 {
		return nil
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rng.Int()
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	maxInt := data[0]

	for _, v := range data {
		maxInt = max(maxInt, v)

	}
	return maxInt
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь

	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup

	maxSlice := make([]int, CHUNKS)
	sizeChunkSlice := len(data) / CHUNKS
	for i := 0; i < CHUNKS; i++ {
		beginIndex := sizeChunkSlice * i
		endIndex := beginIndex + sizeChunkSlice

		if i == CHUNKS-1 {
			endIndex = len(data)
		}

		if beginIndex >= len(data) {
			break
		}

		wg.Add(1)

		go func(idx int, chunk []int) {
			defer wg.Done()

			if len(chunk) == 0 {
				return
			}

			maxInt := chunk[0]
			for _, v := range chunk {
				maxInt = max(maxInt, v)
			}

			maxSlice[idx] = maxInt
		}(i, data[beginIndex:endIndex])
	}

	wg.Wait()

	maxInt := maxSlice[0]
	for _, v := range maxSlice {
		maxInt = max(maxInt, v)
	}

	return maxInt
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	randSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	maxInt := maximum(randSlice)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxInt, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	maxInt = maxChunks(randSlice)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxInt, elapsed.Microseconds())
}
