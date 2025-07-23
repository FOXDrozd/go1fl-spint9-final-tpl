package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t * testing.T){
	const SIZE_5 = 5
	dataTest := generateRandomElements(SIZE_5)
	assert.Len(t, dataTest, SIZE_5)
	
	const SIZE_ZERO = 0
	dataTest = generateRandomElements(SIZE_ZERO)
	assert.Len(t,dataTest, SIZE_ZERO)
}

func TestMaximum(t * testing.T){
 	dataEmpty := make([]int, 0)
	max := maximum(dataEmpty)
	assert.Equal(t,max, 0)
	
 	dataEmpty = []int{1, 2, 3, 4, 5, 9} 
	max = maximum(dataEmpty)
	assert.Equal(t, max, 9)
}
