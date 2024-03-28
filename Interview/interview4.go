package main

import "fmt"

func main() {

	// input := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	// output := []int{}
	// tmp := []int{}
	// max := -1000000000
	// currentMax := 0
	// start := -1

	// for i := 0; i < len(input); i++ {
	// 	currentMax = currentMax + input[i]
	// 	tmp = append(tmp, input[i])
	// 	if currentMax > max {
	// 		max = currentMax
	// 		output = tmp
	// 		start = i
	// 	} else {
	// 		start = -1
	// 	}

	// 	// if currentMax < 0 {
	// 	// 	currentMax = 0
	// 	// 	// start = -1
	// 	// }
	// }

	// fmt.Println(max, output, start)

	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())

}

func intSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func longestCommonPrefix(strs []string) string {
	output := ""
	first := strs[0]
	k := 0
	for i := 1; i < len(strs); i++ {
		for j := 1; j < len(strs); j++ {
			if first[k] != strs[j][k] {
				return output
			}

		}
		k++

		// 00,10,20
		// 01,11,21
		// 02,12,22
	}
	return output
}
