package selectionsortstv

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	unsortedData := []int{0, 4, 5, 1, 6, 23, 4, 2}
	expectedResult := []int{0, 1, 2, 4, 4, 5, 6, 23}
	result := SelectionSort(unsortedData, "asc")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("Test Case Invalid")
	}

	t.Log("Test Case Valid")
}

func ExampleSelectionSort() {
	SelectionSort([]int{2, 1, 4, 6, 7, 0, 8}, "asc")
	// Output : []int{0,1,2,4,6,7,8}
}

func benchmarkSelectionSort(n []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(n, "")
	}
}

func BenchmarkSelectionSort2(b *testing.B) { benchmarkSelectionSort([]int{5, 2}, b) }
func BenchmarkSelectionSort3(b *testing.B) { benchmarkSelectionSort([]int{5, 2, 3}, b) }
func BenchmarkSelectionSort4(b *testing.B) { benchmarkSelectionSort([]int{5, 2, 10, 12}, b) }
func BenchmarkSelectionSort5(b *testing.B) { benchmarkSelectionSort([]int{5, 2, 4, 1, 2}, b) }
