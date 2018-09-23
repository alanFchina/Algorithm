package 剑指offer

func MinNumberRecursion(arr []int) int {
	length := len(arr)
	if length == 1 {
		return arr[0]
	}

	if arr[0] < arr[length-1] {
		return arr[0]
	}

	if length == 2 {
		return MinInOrder(arr)
	}

	mid := length / 2
	if arr[0] == arr[mid] && arr[mid] == arr[length-1] {
		return MinInOrder(arr)
	}

	if arr[mid] >= arr[0] {
		return MinNumberRecursion(arr[mid:])
	} else {
		return MinNumberRecursion(arr[:mid+1])
	}
}
