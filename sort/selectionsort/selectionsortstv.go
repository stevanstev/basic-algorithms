package selectionsortstv

//SelectionSortStv is a sort function for slice of int using selection sort algorithm
func SelectionSortStv(n []int, sortType string) []int {
	var result []int
	var validateSortType bool
	temp := 0
	sliceLen := len(n)
	selectedIndex := 0

	for i := 0; i < sliceLen; i++ {
		sizeSlice := len(n)
		for j := 0; j < sizeSlice; j++ {
			//If Ascending
			validateSortType = temp < n[j]

			//If Descending
			if sortType != "desc" {
				validateSortType = temp > n[j]
			}

			//If start index
			if j == 0 || validateSortType {
				temp = n[j]
				//Get the index of selected number
				selectedIndex = j
				continue
			}

		}

		//Remove the smallest number in slice
		n = append(n[:selectedIndex], n[selectedIndex+1:]...)
		result = append(result, temp)
	}

	return result
}
