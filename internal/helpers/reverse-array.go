package helpers

func ReverseArray(arr []string) []string {
	var newArr []string

	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, arr[i])
	}

	return newArr
}
