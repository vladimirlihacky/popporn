package utils

import "fmt"

func SerializeIntSlice(slice []int) string {
	return fmt.Sprintf("%v", slice)
}

func DeserializeIntSlice(data string) []int {
	var slice []int
	fmt.Sscanf(data, "%v", &slice)
	return slice
}
