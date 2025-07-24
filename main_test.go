package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t * testing.T){
	t.Run("Size = 0", func(t *testing.T){
		const SIZE_ZERO = 0
		dataTest := generateRandomElements(SIZE_ZERO)
		assert.Len(t,dataTest, SIZE_ZERO)
	})
	t.Run("Under different size conditions", func(t *testing.T){
		testCases := []int{1, 200, 400, 20_000, 40_000, 200_000, 400_000}
		for _, size := range testCases {
			dataTest := generateRandomElements(size)
			assert.Len(t, dataTest, size)
		}
	})
}

func TestMaximum(t * testing.T){
	t.Run("Empty data", func(t *testing.T){
		dataEmpty := make([]int, 0)
		max := maximum(dataEmpty)
		assert.Equal(t,max, 0)
	})
	
	t.Run("Test at different values", func(t *testing.T){
		testCases := []struct{
			data []int
			max int
		}{
			{
				data: 	[]int{1, 2, 3, 4, 5, 9},
				max: 	9,
			},
			{
				data: 	[]int{100, 200, 300, 400, 500, 900},
				max: 	900,
			},
					{
				data: 	[]int{100_000, 200_000, 300_000, 400_000, 900_000, 500_000},
				max: 	900_000,
			},
		}

		for _, val := range testCases {
			max := maximum(val.data)
			assert.Equal(t, max, val.max)
		}
	})
}

func TestMaxChunk(t * testing.T){
	t.Run("Check correction work", func(t *testing.T){
		testCases := []struct{
	
			data []int
			max int
		}{
			{
				data: 	[]int{1, 2, 3, 4, 5, 9},
				max: 9,
			},
			{
			
				data: 	[]int{
					100, 200, 300, 400, 500, 900, 140, 250, 
					350, 450, 550, 950, 600, 550, 1350, 850,
				 	250, 750, 450, 550, 950, 600, 550, 850,
					300, 400, 500, 900, 140, 250, 300, 400, 
					500, 900, 140, 250, 300, 400, 500, 900, 
					140, 250, 300, 400, 500, 900, 140, 250, 
					300, 400, 500, 900, 140, 250, 300, 400, 
					500, 900, 140, 250, 300, 400, 500, 900, 
					140, 250, 300, 400, 500, 900, 140, 250, 
				},
				max: 	1350,
			},
			{
		
				data: 	[]int{100_000, 200_000, 300_000, 400_000, 900_000, 500_000, 100_000, 200_000, 300_000, 400_000, },
				max: 	900_000,
			},
		}

		for _, val := range testCases {
			max := maxChunks(val.data);
			assert.Equal(t, max , val.max)
		}
	})
}