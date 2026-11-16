package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size == 0 {
		return nil
	}
	rng := rand.NewSource(int64(size))
	var slice []int
	for i := 0; i < size; i++ {
		slice = append(slice, int(rng.Int63()))
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if data == nil {
		return 0
	}

	var maxInt int

	for _, v := range data {
		maxInt = max(maxInt, v)

	}

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	var wg sync.WaitGroup

	wg.Add(CHUNKS)

	var maxSlice []int
	for i := 0; i <= CHUNKS; i++ {
		sizeChunkSlice := len(data) / CHUNKS

	}
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
