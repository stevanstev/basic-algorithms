package bubblesortstv

import (
	"reflect"
	"testing"
)

func TestBubbleSortStv(t *testing.T) {
	unsortedData := []int{0, 4, 5, 1, 6, 23, 4, 2}
	expectedResult := []int{0, 1, 2, 4, 4, 5, 6, 23}
	result := bubbleSortStv(unsortedData, "asc")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("Test Case Invalid")
	}

	t.Log("Test Case Valid")
}

func ExamplebubbleSortStv() {
	bubbleSortStv([]int{2, 1, 4, 6, 7, 0, 8}, "asc")
	// Output : []int{0,1,2,4,6,7,8}
}

func benchmarkBubbleSortStv(n []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSortStv(n, "")
	}
}

func BenchmarkBubbleSort2(b *testing.B) { benchmarkBubbleSortStv([]int{5, 2}, b) }
func BenchmarkBubbleSort3(b *testing.B) { benchmarkBubbleSortStv([]int{5, 2, 3}, b) }
func BenchmarkBubbleSort4(b *testing.B) { benchmarkBubbleSortStv([]int{5, 2, 10, 12}, b) }
func BenchmarkBubbleSort5(b *testing.B) { benchmarkBubbleSortStv([]int{5, 2, 4, 1, 2}, b) }
