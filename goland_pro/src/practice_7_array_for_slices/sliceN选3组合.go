package main

import "fmt"

func main() {
	var a = []int32{11, 22, 33, 44, 55, 66}
	fmt.Println(mix(0, 1, 2, a, nil))
	// fmt.Println(max(a))

}

func mix(i, j, k int, a []int32, result [][]int32) [][]int32 {
	length := len(a)
	if k <= length-1 {
		result = append(result, []int32{a[i], a[j], a[k]})
		k++
		result = mix(i, j, k, a, result)
	} else {
		if j <= length-2 {
			j++
			k = j + 1
			result = mix(i, j, k, a, result)
		} else {
			if i <= length-3 {
				i++
				j = i + 1
				k = j + 1
				result = mix(i, j, k, a, result)
			}
		}
	}
	return result
}
