package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	temp := map[int]struct{}{}
	for _, v := range arr {
		if _, ok := temp[v]; !ok {
			temp[v] = struct{}{}
			result = append(result, v)
		}
	}

	fmt.Println(result)
}
