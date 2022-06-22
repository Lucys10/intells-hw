package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	var min, max int32

	r := strings.Split(input, " ")

	numbInt32 := make([]int32, 0, len(r))
	for _, v := range r {
		vInt, _ := strconv.Atoi(v)
		numbInt32 = append(numbInt32, int32(vInt))
	}

	if len(numbInt32) == 1 {
		result = strconv.Itoa(int(numbInt32[0]))
		fmt.Println(result)
		return
	}

	min = numbInt32[0]
	max = numbInt32[0]

	for _, v := range numbInt32 {

		switch {
		case v < min:
			min = v
		case v > max:
			max = v
		}
	}

	maxMin := []string{strconv.Itoa(int(max)), strconv.Itoa(int(min))}
	result = strings.Join(maxMin, " ")
	fmt.Println(result)

}
