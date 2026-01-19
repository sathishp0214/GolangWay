package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {

	// s := 10
	// h := strconv.Itoa(s)
	// fmt.Printf("%T, %s, %v", h, h, h)

	// d := "sathish is good and nice and brainy"

	// f := []byte(d)
	// fmt.Println(f, []rune(d))

	// g := []byte{115, 97, 116}
	// g1 := 115
	// fmt.Println(string(g), string(g1), d[len(d)-1])

	// df := make([]string, len(d))
	// df = append(df, "dd")
	// fmt.Println(df, len(df))

	// findingVowels()

	// frequencyCharacters()

	// ReplaceCapitalStartingLetter()

	// RemoveDuplicateCharacters()

	// fmt.Println("RemoveDuplicateGenerics -- ", RemoveDuplicateGenerics([]int{2, 2, 44, 5, 6, 7, 7, 8, 8}))
	// fmt.Println("RemoveDuplicateGenerics -- ", RemoveDuplicateGenerics([]rune{65, 65, 44, 5, 6, 97, 97, 8, 8}))
	// fmt.Println("RemoveDuplicateGenerics -- ", RemoveDuplicateGenerics([]string{"s", "d", "s", "e", "e"}))
	// fmt.Println("RemoveDuplicateGenerics -- ", RemoveDuplicateGenerics([]bool{true, false, true}))

	// secondMinimum()

	// fmt.Println("anagrams -- ", anagramString("listen", "silent"))
	// fmt.Println("anagrams -- ", anagramString("boss", "sobb")) //this should have returned false

	// factorialNumber(5)
	// countFrequency("sathish have a nice day")

	// fmt.Println("reverse string---:", reverseString1("sathish have a nice day"))

	// a := 10
	// b := 20
	// a = a + b
	// b = a - b // b 10
	// a = a - b // a 20

	// fmt.Println(a, b)

	// s := []int{1, 4, 5, 8, 10, 12, 14, 16, 20, 25}
	// missingNumbers := []int{}
	// for i := 0; i < len(s)-1; i++ {
	// 	if s[i+1]-s[i] > 1 {
	// 		for j := s[i] + 1; j < s[i+1]; j++ {
	// 			// fmt.Println("missing numbers,", j)
	// 			missingNumbers = append(missingNumbers, j)

	// 		}
	// 	}
	// }

	// fmt.Println("missing numbers in the slice - ", missingNumbers)

	// fmt.Println("Find missing numbers in slice - ", FindMissingNumbers([]int{1, 4, 5, 8, 10, 12, 14, 16, 20, 25}))

	// fmt.Println("Find missing numbers in slice - ", findDuplicates([]int{1, 2, 3, 2, 4, 5, 1, 10, 3}))

	// nums1 := []int{2, 11, 15, 7}
	// target1 := 9
	// result1 := twoSum(nums1, target1)
	// fmt.Printf("Input: %v, Target: %d, Result1: %v\n", nums1, target1, result1)

	// a := "sathish"
	// b := a
	// b = "sat"

	// fmt.Println(&a, &b, a, b)

	// result := make([]int, 5)
	// fmt.Println(result, result[1] == 0)

	// c := 'z'
	// fmt.Println(reflect.TypeOf(c), c)

	// var p3 chan interface{}
	// fmt.Println(p3, reflect.TypeOf(p3))

	// a := 123
	// n := 0
	// for a > 0 {
	// 	t := a % 10
	// 	n = n*10 + t
	// 	a = a / 10
	// }

	// fmt.Println(a, n, 1%10, 8%10, 1/10, 8/10)

	// a := 20
	// b := 70

	// for i := a; i <= b; i++ {
	// 	flag := false
	// 	for j := 2; j < i; j++ {
	// 		if i%j == 0 {
	// 			flag = true
	// 			break
	// 		}
	// 	}

	// 	if !flag {
	// 		fmt.Println("prime number ---", i)
	// 	}
	// }

	// s := "Golang is simple but powerful  is    aw"
	// count := 0
	// for i := 0; i < len(s); i++ {
	// 	if string(s[i]) == " " && string(s[i+1]) != " " { //counts the space to count the words
	// 		count++
	// 	}
	// }

	// fmt.Println("total words in the line - ", count+1)

	// m1 := [][]int{{1, 2}, {3, 4}}
	// m2 := [][]int{{5, 6}, {7, 8}}
	// m3 := make([][]int, len(m1))

	// fmt.Println(m3, len(m3), m1, m2)

	// for i := 0; i < len(m1); i++ {
	// 	m3[i] = make([]int, len(m1))
	// 	for j := 0; j < len(m1); j++ {
	// 		m3[i][j] = m1[i][j] + m2[i][j]
	// 	}
	// }

	// fmt.Println("sum of matrix-----", m3)

	// e := 23.5

	// fmt.Println(e > 100, -70 > -80, -70 < -80, -30 > -35, -40 > -35)

	// // now := time.Now()

	// // // Calculate yesterday's date using AddDate (recommended over subtracting 24h due to DST)
	// // yesterday := now.AddDate(0, 0, -2)

	// // // Extract the day of the month as a number (int)

	// // dayNumber := yesterday.Day()

	// // fmt.Println("Yesterday's date number is: %d\n", dayNumber)

	// // (true || true) || (true || true)

	// if (false && true) || ((true && false) || true) || ((false && false) || false) {
	// 	fmt.Println("passed-------")
	// }

	// fmt.Println(123 < 35)

	// var a int64
	// a = 100

	// fmt.Println(a / 3)

	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(s[1:])

	// loc, _ := time.LoadLocation("Asia/Kolkata")

	// 1️⃣ Compute first aligned run
	// firstRun := nextAlignedTime(5*time.Minute, 5*time.Second, loc)

	// fmt.Println("firstrun returns-------", firstRun, time.Until(firstRun))

	// time.Sleep(time.Until(firstRun))

	// now := time.Now().In(loc)
	// fmt.Println("task started -------", firstRun, now)

	// // 2️⃣ Wait until first aligned execution
	// timer := time.NewTimer(time.Until(firstRun))

	// fmt.Println(firstRun, "-------------", <-timer.C)
	val := 340.0
	result := (50.0 / 100.0) * val
	fmt.Println(result) // ✅ "350"
	fmt.Println(math.Trunc(7.859999999999999))

	fmt.Println(Truncate2(7.859999999999999)) // 7.85
	fmt.Println(Truncate2(6.134))
	fmt.Println(Truncate2(6.0))
	fmt.Println(Truncate2(6.4))
	fmt.Println(Truncate2(6.8))
	fmt.Println(ReversalOption("NIFTY20JAN26C25800"), len("NIFTY20JAN26C25800"))
	// lastPriceInt := 2437
	// strikePrice := 1
	// // lastPriceInt = 787
	// // strikePrice = 5
	// fmt.Println(lastPriceInt / strikePrice)
	// lastPriceInt = (lastPriceInt / strikePrice) * strikePrice
	// fmt.Println(lastPriceInt)
	// lastPriceInt = lastPriceInt + strikePrice
	// fmt.Println(lastPriceInt)

	// m := 0 + 10
	// for i := 0; i < m; i += 3 {
	// 	fmt.Println(i, 100.50/5)
	// }
	// a, b := strconv.Atoi("245.56")
	// fmt.Println(a, b)

}

// "currently allows for nifty index options only - EX: "NIFTY20JAN26P25800"
func ReversalOption(symbol string) string {
	n := len(symbol)

	if !strings.HasPrefix(symbol, "NIFTY") || n != 18 {
		fmt.Println("currently allows for nifty index options only----------")
		return ""
	}

	if symbol[n-6] == 'P' {
		return symbol[:n-6] + "C" + symbol[n-5:]
	}
	if symbol[n-6] == 'C' {
		return symbol[:n-6] + "P" + symbol[n-5:]
	}
	return symbol
}

func Truncate2(val float64) float64 {
	return math.Trunc(val*100) / 100
}

func nextAlignedTime(interval time.Duration, buffer time.Duration, loc *time.Location) time.Time {
	now := time.Now().In(loc)

	elapsed := now.Truncate(interval)
	next := elapsed.Add(interval)

	fmt.Println(now, "--------", elapsed, "-------------", next)

	return next.Add(buffer)
}

func dummy() {
	fmt.Println("dummy function Entered---------")
	time.Sleep(10 * time.Second)
}

// "{[()]}"
func isValidBrackets(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // Create a hash map to store numbers and their indices

	for i, num := range nums {
		complement := target - num // Calculate the complement needed to reach the target

		// Check if the complement exists in the map
		if index, found := numMap[complement]; found {
			return []int{index, i} // If found, return the indices
		}

		// If not found, add the current number and its index to the map
		numMap[num] = i

		fmt.Println("ssssssssssss", i, num, complement, numMap)
	}

	return nil // If no solution is found, return nil
}

// func twoSum(n []int, target int) (int, int) {
// 	if len(n) < 2 {
// 		return -1, -1
// 	}

// 	first := n[0]
// 	for i := 1; i < len(n); i++ {
// 		if first+n[i] == target {
// 			return first, n[i]
// 		}

// 		first = n[i]
// 	}

// 	return -1, -1
// }

func FindMissingNumbers(s []int) []int {
	if len(s) < 2 {
		return []int{}
	}

	presentMap := map[int]bool{}
	missingNumbers := []int{}
	start := s[0]
	end := s[len(s)-1]

	for _, i := range s {
		presentMap[i] = true
	}

	for i := start; i <= end; i++ {
		if !presentMap[i] {
			missingNumbers = append(missingNumbers, i)
		}
	}

	return missingNumbers

}

func findDuplicates(nums []int) []int {
	m := make(map[int]int)
	res := []int{}
	for _, n := range nums {
		m[n]++
	}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

func reverseString1(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
		fmt.Println(i, j)
	}
	return string(runes)
}

// simplied version
func countFrequency(v string) {
	output := map[string]int{}
	for _, i := range v {
		output[string(i)]++
	}
	fmt.Println("frequency of characters-----", output)
}

func factorialNumber(num int) {
	output := 1
	for i := 1; i <= num; i++ {
		output = output * i
	}

	fmt.Println("factorial output -- ", output)
}

// func insertIntoSliceGenerics[t int | string](s []t, value t, index int) []t {
// 	if index > len(s) {
// 		return nil
// 	}

// 	if index == len(s) {
// 		s = append(s, value)
// 		return s
// 	}

// 	tmp := s[:index+1]

// }

// not correct logic for tricky anagrams
func anagramString(w, w1 string) bool {

	if len(w) != len(w1) {
		return false
	}

	for _, i := range w {
		flag := false
		for _, j := range w1 {
			if i == j {
				flag = true
				break
			}
		}

		if !flag {
			return false
		}

	}

	return true
}

func secondMinimum() {
	k := []int{-12, 0, 100, 2, 3, -4, 5, -1, 10, -7, -20}
	minOne := k[0]
	minTwo := k[1]

	if minOne > minTwo {
		minOne, minTwo = minTwo, minOne
	}

	for i := 2; i < len(k); i++ {
		if k[i] >= minTwo {
			continue
		}

		if k[i] < minOne {
			minTwo = minOne
			minOne = k[i]
		}

		if k[i] > minOne && k[i] < minTwo {
			minTwo = k[i]
		}

	}

	fmt.Println("second minimum one", minOne, minTwo)
}

func minimumMaximum() {
	k := []int{0, 100, 2, 3, -4, 5, -1, 10}
	max := k[0]
	min := k[0]
	for i := 1; i < len(k); i++ {
		if k[i] < min {
			min = k[i]
		}

		if k[i] > max {
			max = k[i]
		}
	}

	fmt.Println(min, max)
}

func ReplaceCapitalStartingLetter() {
	d := "sathish is good and nice and brain"
	words := strings.Split(d, " ")
	for index, i := range words {
		tmp := []byte(i)
		tmp[0] = tmp[0] - 32
		words[index] = string(tmp)
	}

	fmt.Println("capital starting letter - ", strings.Join(words, " "))
}

func RemoveDuplicateCharacters() {
	d := "sathishhellohowareyou"
	charMap := map[rune]bool{}
	output_string := ""

	for _, i := range d {
		if _, exists := charMap[i]; !exists {
			charMap[i] = true
			output_string = output_string + string(i)

		}
	}

	fmt.Println("freq characters", charMap, output_string)
}

// generics good use case
func RemoveDuplicateGenerics[generics int | string | rune | bool](s []generics) []generics {
	tmp := map[generics]bool{}
	output := []generics{}

	for _, i := range s {
		if _, exists := tmp[i]; !exists {
			tmp[i] = true
			output = append(output, i)
		}
	}

	return output
}

func frequencyCharacters() {
	d := "sathish is good and nice and brain"
	charMap := map[string]int{}
	for _, i := range d {
		// if charMap[string(i)] == 0 {
		if _, exists := charMap[string(i)]; !exists {
			charMap[string(i)] = 1

		} else {
			charMap[string(i)]++
		}
	}

	fmt.Println("freq characters", charMap)
}

func findingVowels() {
	d := "sathish is good and nice and brain"
	vowels := "aeiou"
	var vowelsCount int
	var nonVowelsCount int
	for _, i := range d {

		if string(i) == " " {
			continue
		}
		// if strings.Contains(vowels, string(i)) {
		if IsContains(vowels, string(i)) {
			vowelsCount++
		} else {
			nonVowelsCount++
		}
	}
	fmt.Println("vowels count", vowelsCount, nonVowelsCount)
}

func IsContains(substr string, s string) bool {
	for _, i := range substr {
		if string(i) == s {
			return true
		}
	}
	return false

}

func reverseStringUsingByte() {
	d := "sathish"

	// e := []byte{}
	// for i := len(d) - 1; i >= 0; i-- {
	// 	e = append(e, d[i])
	// }

	//another way
	e := make([]byte, len(d))
	j := 0
	for i := len(d) - 1; i >= 0; i-- {
		e[j] = d[i]
		j++
	}
	fmt.Println(string(e), len(e))

	//another way
	d1 := []byte(d)
	k := len(d1) - 1
	for i := 0; i < len(d1)/2; i++ {
		d1[i], d1[k] = d1[k], d1[i]
		k--
	}
	fmt.Println(string(d1))
}
