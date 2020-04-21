package bubblesortstv

//BubbleSort is a sort function for slice of int using bubble sort algorithm
func BubbleSort(n []int, sortType string) []int {
	sliceLen := len(n)

	//4, 5, 1, 2, 3
	for i := 0; i < sliceLen; i++ {
		isSorted := 1
		for j := 0; j < sliceLen-1; j++ {
			validateType := n[j] > n[j+1]
			if sortType == "desc" {
				validateType = n[j] < n[j+1]
			}

			if validateType {
				temp := n[j]
				n[j] = n[j+1]
				n[j+1] = temp

				//if there's swapping
				isSorted = 0
			}
		}

		//if all slice are sorted, then break
		if isSorted == 1 {
			break
		}
	}

	return n
}
