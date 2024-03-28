package main

import (
	"fmt"
	"sync"
)

var pr = fmt.Println

func main() {

	// secondMaximumWithoutSorting()

}

// this executing, But this matrix multiplication is wrong mathmatically
func MatrixMultiplication() {
	m1 := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	m2 := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	result := make([][]int, len(m1))
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1[0]); j++ {
			result[i] = append(result[i], m1[i][j]*m2[i][j])
		}
	}

	fmt.Println("matrix multiplication", result)

}

func permuteIterative(str string) {
	n := len(str)

	stack := make([]int, n)
	for i := range stack {
		stack[i] = 0
	}

	fmt.Println("Permutations:")
	fmt.Println(str)

	i := 0
	for i < n {
		if stack[i] < i {
			if i%2 == 0 {
				str = swap(str, 0, i)
			} else {
				str = swap(str, stack[i], i)
			}
			fmt.Println(str, stack[i], "=======", i)
			stack[i]++
			i = 0
		} else {
			stack[i] = 0
			i++
		}
	}
}

func swap(str string, i, j int) string {
	strBytes := []byte(str)
	strBytes[i], strBytes[j] = strBytes[j], strBytes[i]
	return string(strBytes)
}

func sampleTask(wg *sync.WaitGroup) {
	for i := 0; i < 10000000; i++ {
		_ = 10
	}
	wg.Done()
}

func task2() {
	pr("hello")
}

func SwapCharactersForStringsEqual() {
	s1 := "bank"
	s2 := "kanb"
	tmp := []byte(s2)
	pr(tmp, s2)
	for i := 0; i < len(s1); i++ {

	}

}

func longestPalindromeNumber() {
	s := "forgeeksskeegfor"
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			fmt.Println(string(s[i]), "===", string(s[j]))
		}
	}
}

func longestPalindrome() string {
	s := "forgeeksskeegfor"
	start, maxLen, l := 0, 0, len(s)
	for i := 0; i < l; i++ {
		expandAroundCenter(s, i, i, &start, &maxLen)
		expandAroundCenter(s, i, i+1, &start, &maxLen)
	}
	return s[start : start+maxLen]
}

func expandAroundCenter(s string, l, r int, start, maxLen *int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l, r = l-1, r+1
	}
	if r-l-1 > *maxLen {
		*maxLen = r - l - 1
		*start = l + 1
	}
}
