func findRightMax(rightMax int, element int) int {
	if element > rightMax {
		rightMax = element
	}
	return rightMax
}

func replaceElements(arr []int) []int {
	rightMax := -1
	for i := len(arr) - 1; i >= 0; i-- {
		current := arr[i]
		arr[i] = rightMax
		rightMax = findRightMax(rightMax, current)
	}
	return arr
}