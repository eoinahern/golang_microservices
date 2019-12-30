package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	elms := []int{9, 8, 7, 6, 5}
	sorted := BubbleSort(elms)

	assert.NotNil(t, sorted)
	assert.EqualValues(t, 5, len(elms))
	assert.EqualValues(t, 5, len(sorted))
	assert.EqualValues(t, 5, sorted[0])
	assert.EqualValues(t, 6, sorted[1])
}

func getList(n int) []int {

	nums := make([]int, n)
	x := 0
	for i := n; i >= 1; i-- {
		nums[x] = i
		x++
	}

	return nums
}

func BenchmarkGetStuff(b *testing.B) {

	for i := 0; i < b.N; i++ {
		BubbleSort(getList(10))
	}

}
