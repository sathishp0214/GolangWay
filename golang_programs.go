package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	_ "reflect"
	"sort"
	"strconv"
	_ "strconv" // "_" avoids "not used compile errors"
	"strings"
	"time"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var romanMap = map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M"}

// romanMap := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XL": 40, "L": 50}

var romanLettersInDescending = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

func main() {

	// rt := []string{"don", "donkey", "donk"}
	// out := ""
	// flag := 0
	// for i := 0; i < len(rt); i++ {
	// 	flag = 0
	// 	tmp := rt[i][0]
	// 	for j := i + 1; j < len(rt); j++ {
	// 		if rt[j] != string(tmp) {
	// 			flag = 1
	// 		}
	// 	}
	// 	if flag == 0 {
	// 		out = out + string(tmp)
	// 		fmt.Println("+++++++", out, string(tmp))
	// 	}

	// }

	//prints number 1 to 100 with loop, using recursion
	// printnumbers(100)

	//finding longest consective event numbers in slice
	// q := []int{1, 2, 4, 6, 7, 8}
	// q := []int{3, 6, 7, 2, 4, 6, 9, 10, 12, 14, 16}
	// result := []int{}
	// // k := 3
	// g := 0
	// for _, j := range q {
	// 	if j%2 == 0 {
	// 		g = g + 1
	// 	} else {
	// 		result = append(result, g)
	// 		g = 0
	// 	}
	// }
	// //adding final set of g value
	// result = append(result, g)

	// fmt.Println(g, result, maximumInt(result))

	//rotate slice elements using slicing
	// w := []int{11, 12, 13, 14, 15, 16}
	// rotate := 5
	// p := []int{}
	// p = append(p, w[rotate:]...)
	// p = append(p, w[:rotate]...)
	// fmt.Println(p)

	//need to check this, value assigning
	// r := []int{2, 3, 4, 5, 6}
	// t := r[:2]
	// t = append(t, 56)
	// fmt.Printf("%p %p", &r, &t)
	// fmt.Println(r, t)

	//finding matching substring between all values in slice
	// st1 := []string{"rud", "rudra", "rahi", "ghh", "rughr", "rudp", "rahikl"}
	// for i := 0; i < len(st1); i++ {

	// 	for j := i + 1; j < len(st1); j++ {
	// 		if strings.Contains(st1[j], st1[i]) {
	// 			fmt.Println(st1[j], st1[i])
	// 		}
	// 	}

	// }

	// fmt.Println(ValidateEventInt64Type(2332362477.466))

	// fmt.Println(findSubString("sathish", "athi"))
	// fmt.Println(findSubStringAllOccurences("sathish hello athi yesathi", "athi"))

	// missingNumbersInSortedSlice([]int{6, 7, 10, 11, 13})

	birthdayCakeCandles([]int{3, 2, 1, 3})

	fmt.Println(timeConversion12Hourto24Hour("12:34:50PM")) //12:34:50
	fmt.Println(timeConversion12Hourto24Hour("07:34:50PM")) //19:34:50

	sortMapsByValue()

	LargestConsectiveOccurence()

	countValuesOccurancesOnslice()

	equalSubsetSum()

	// sum_of_digits(123)
	// number_of_digits(7869)
	// reverseNumber(1234)

	// countWords()

	// reverseString()

	// primeNumbers(100)

	// perfectNumber(28)

	// intstringslicesumOfDigits()

	swappedWithoutTemporaryValue()
	staircaseReversePrintPattern(7)
	bubbleSort()

}

func swappedWithoutTemporaryValue() {
	a := 10
	b := 20
	a = a + b
	b = a - b //b gets a's swapped value
	a = a - b //a gets a's swapped value
	fmt.Println("swapped valued without temporary value")
}

// func IsValidPassword() {
// 	regexp.MustCompile(str)
// }

func sortMapsByValue() {
	// Write your code here
	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4, 3}
	er := map[int32]int32{}
	sliceKeys := []int32{}
	for _, i := range arr {
		if _, exists := er[i]; !exists {
			er[i] = 1
		} else {
			er[i] = er[i] + 1
		}
	}

	for key, _ := range er {
		sliceKeys = append(sliceKeys, key) //slicekeys has all maps keys
	}

	//this reverse sorts map by values
	sort.SliceStable(sliceKeys, func(i, j int) bool {
		return er[sliceKeys[i]] > er[sliceKeys[j]] // map[slice[keyValue]]
	})

	fmt.Println(sliceKeys, "--------", er)
	fmt.Println("sorted map by maximum value --", sliceKeys[0]) //input slice has maximum 3 value occurance
}

func timeConversion12Hourto24Hour(s string) string {

	var PMFlag bool
	if strings.HasSuffix(s, "PM") {
		PMFlag = true
		s = strings.TrimSuffix(s, "PM") //24 hour not shows AM or PM value, so trims that
	} else {
		s = strings.TrimSuffix(s, "AM")
	}

	n := strings.Split(s, ":")

	hour, _ := strconv.Atoi(n[0])

	if PMFlag && hour >= 1 && hour <= 11 { //from 1 pm to 11 pm adds 12 to hour for 24 hrs time
		hour = hour + 12
		n[0] = strconv.Itoa(hour)
		r := strings.Join(n, ":")
		return r
	}

	if !PMFlag && hour == 12 { //if AM time, then its 12AM something then changes hour into 00
		n[0] = "00"
		r := strings.Join(n, ":")
		return r
	}

	return s

}

func birthdayCakeCandles(candles []int) int { //birthdayCakeCandles([]int{3, 2, 1, 3})
	// Write your code here
	sort.SliceStable(candles, func(i, j int) bool {
		return candles[i] > candles[j]
	})

	fmt.Println(candles)
	count := 1
	for i := 0; i < len(candles)-1; i++ {
		if candles[i] != candles[i+1] {
			return count
		}
		count++
	}

	return count
}

func staircaseReversePrintPattern(n int32) {
	// 	  	 #
	// 	 	##
	//     ###
	//    ####
	//   #####
	//  ######
	// #######
	d := int(n)
	for i := d; i >= 1; i-- {
		for j := 1; j <= d; j++ {
			if j >= i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func equalSubsetSum() {
	// inp := []int{1, 5, 11, 5, 3, 3}
	inp := []int{5, 10, 5}
	sum := 0
	sort.SliceStable(inp, func(i, j int) bool {
		return inp[i] > inp[j] //reverse sort for check from 0th index
	})

	for _, i := range inp {
		sum = sum + i
	}

	OffValue := sum / 2 //if float value got, output differs
	tmpSum := 0
	for i := 0; i < len(inp)-1; i++ {
		tmpSum = tmpSum + inp[i]
		if tmpSum > OffValue {
			fmt.Println("equally split slice values not found here")
			break
		}
		if tmpSum == OffValue {
			fmt.Println(inp, "equally split slice values", inp[:i+1], "---", inp[i+1:]) //{1, 5, 11, 5, 3, 3}  ---> [11 3] --- [5 5 3 1]
			break
		}

		diff := OffValue - tmpSum
		// tmpInp := inp[i+1:]
		valueIndex := slices.Index(inp, diff) //checks the expected difference value index in slice and swaps to next index, So the next index for loop iteration takes that.
		fmt.Println(inp, "----", diff, tmpSum, i, valueIndex)
		//TODO -- Having limitation, for example if diff 6, but the inp slice has 3,3, Still index returns -1, So need a fix
		if valueIndex < 0 {
			fmt.Println("equally split slice values not found")
			break
		}
		inp[i+1], inp[valueIndex] = inp[valueIndex], inp[i+1]

	}

}

func longestPalindromeSubstring() {
	// s := "satmadampp"
}

func countValuesOccurancesOnslice() {
	out := []int{}
	inp := []int{1, 1, 2, 2, 2, 3, 4, 5, 5, 7, 7, 7, 8, 1, 1, 1}
	count := 1
	for i := 0; i < len(inp)-1; i++ {
		if inp[i] == inp[i+1] {
			count++
		} else {
			out = append(out, count)
			count = 1
		}
	}

	//appends the last count value
	out = append(out, count)

	fmt.Println("countValuesOccurancesOnslice", out) //[2 3 1 1 2 3 1 3]
}

func CircularLoopIteration() {
	//circular loop iteration
	size := 7
	startingPosition := 3
	for i := 0; i < size; i++ {
		index := (i + startingPosition) % size
		fmt.Println(index, " ") //3 4 5 6 0 1 2  //starts from startingPosition's 3
	}
}

func LargestConsectiveOccurence() {
	df := []int{2, 3, 4, 4, 5, 6, 6, 6, 8, 8, 8, 1, 1, 6, 6, 6, 6, 6}
	outputMap := map[int]int{}
	count := 1
	for i := 0; i < len(df)-1; i++ {
		if df[i] == df[i+1] {
			count++
			outputMap[df[i]] = count //don't need to check the exiting key in map, because defaulty if key not there it will create key and value, If key is already there it will update the value alone.
		} else {
			count = 1 //whenever there is no consective value making count back to 1 value
		}
	}
	fmt.Println(outputMap)

	fmt.Println("largest consective occurance", slices.Max(maps.Values(outputMap))) // getting map values alone and finding maximum value from slice

}

func fiabonciiUsingLoop() {
	//Fianbioncii series using loop

	a := 0
	b := 1
	c := 0
	fmt.Printf("%d %d", a, b)
	for i := 0; i < 10; i++ {
		c = a + b
		fmt.Printf(" %d ", c)
		a = b
		b = c

	}
}

func removeDuplicateUsingMap() {
	//remove duplicates in slice (using maps) - (For string converts into arrays and checks duplicates)

	input_array := []string{"a", "b", "c", "a", "d", "e", "b", "c", "f", "1", "1", "8", "0"}

	result_array := []string{} //empty dynamic length array //similar in python like "results list for append"

	mid_map := map[string]bool{}

	for i := range input_array {
		// tmp := input_array[i]
		_, y := mid_map[input_array[i]]
		if !y { //if key is not there, creating key and storing values in result_array
			mid_map[input_array[i]] = true
			result_array = append(result_array, input_array[i])
		}
	}

	fmt.Println(result_array) //[a b c d e f 1 8 0]
}

func countOccurences() {
	//Count the letters occurances in maps

	st := "sathish have a nice day"

	data := map[string]int{}

	for i := range st {
		tmp := string(st[i])
		if tmp != " " { //filters spaces
			_, y := data[tmp]
			if y {
				data[tmp] = data[tmp] + 1
			} else {
				data[tmp] = 1
			}
		}
	}

	fmt.Println(data) //map[a:4 c:1 d:1 e:2 h:3 i:2 n:1 s:2 t:1 v:1 y:1]
}

func printPattern() {
	//print star pattern
	// *****
	// ****
	// ***
	// **
	// *

	out1 := 5

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < out1; i++ {
		var j int
		for j = 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func insertSlice() {
	//insert an element in an array using index
	jk := []int{10, 20, 30, 40, 50}

	df := []int{}
	insert_index := 1
	insert_value := 95

	df = append(jk[:insert_index+1], jk[insert_index:]...) //This appends together [10,20], [20,30,40,50]

	df[insert_index] = insert_value // inserts in the index

	fmt.Println(df)
}

func removeSlice() {
	//remove an element in an array using index

	jk := []int{10, 20, 30, 40, 50}

	df := []int{}
	delete_index := 2

	df = append(jk[:delete_index], jk[delete_index+1:]...)
	fmt.Println("after deletion - ", df)
}

func minimumMaximumValues() {
	//minimum and maximum values in an array

	ks := []int{10, 20, 30, 40, 60, 50, 9}

	min := ks[0]
	max := ks[0]

	for i := 1; i < len(ks); i++ {
		if ks[i] < min {
			min = ks[i]
		}
		if ks[i] > max {
			max = ks[i]
		}
	}

	fmt.Println("minimum and maximum value in array - ", min, max)

	// maximum and minimum key (or) maximum and minimum value from map

	d_map := map[int]string{7: "hhhh", 8: "KKKK", 120: "EREER", 9: "PPPPP"}
	max_key := -1
	key_value := ""

	for key, value := range d_map {
		if key > max_key {
			max_key = key
			key_value = value
		}

	}

	fmt.Println(max_key, key_value)
}

func Palindrome() {
	// checks palindrome on string

	palindrome_str := "masdam"
	flag := 0
	limit := len(palindrome_str) - 1
	for i := 0; i <= limit/2; i++ {
		if palindrome_str[i] != palindrome_str[limit-i] {
			flag = 1
			break
		}
	}

	if flag == 0 {
		fmt.Println("It is Palindrome")
	} else {
		fmt.Println("It is not a Palindrome")
	}

}

func reverseStringDifferent() {
	//reverse a string without slice
	st := "sathish"
	reverse := ""

	for _, character := range st {
		reverse = string(character) + reverse
	}

	fmt.Println(reverse)
}

func ReverseSlice() {
	//reverse slice  -- (Like python - , convert string into array and can reverse it)

	xy := []string{"s", "a", "t", "h", "i", "s", "h", " ", "i", "s", " ", "g", "o", "o", "d"}
	// xy := []string {"1","2","3","4","5","6"}
	// xy := []int {10,12,14,16,18}
	limit := len(xy) - 1
	for i := 0; i <= limit/2; i++ {
		xy[i], xy[limit-i] = xy[limit-i], xy[i]
	}

	fmt.Println(xy) //reverse string array //[d o o g   s i   h s i h t a s]
}

func ReverseNumber() {
	//Reverse a number

	num := 12345
	reverse_num := 0

	for num > 0 {
		remainder := num % 10
		reverse_num = reverse_num*10 + remainder
		num = num / 10
	}

	fmt.Println("reverse_number - ", reverse_num)
}

func armstringNumber() {
	//Armstrong number (using int to string to int cnversion)

	num := 371
	sum := 0
	str_num := strconv.Itoa(num)

	for i := 0; i < len(str_num); i++ {
		tmp := str_num[i]
		newInt, _ := strconv.Atoi(string(tmp)) //tmp is in "uint8" datatype, so changes into "string". Then converts string into integer
		sum = sum + (newInt * newInt * newInt)
	}
	fmt.Println("sum - ", sum)

	if num == sum {
		fmt.Println("Armstrong number")
	} else {
		fmt.Println("Not a Armstrong number")
	}

	// Armstrong number using normal logic
	number := 153
	sum1 := 0
	for number > 0 {
		tmp := number % 10
		sum = sum + (tmp * tmp * tmp)
		number = number / 10
	}

	fmt.Println(sum1)
}

func ReverseStringSlice() {
	// 	//reverse a string
	gh := []byte("Sathish is hello")
	j := len(gh) - 1
	for i := 0; i < len(gh)/2; i++ {
		fmt.Println(i, j, len(gh))
		gh[i], gh[j] = gh[j], gh[i]
		j = j - 1
	}
	fmt.Println("reverse string--", string(gh))
}

func LeapYear() {
	//leap year - (good example of nested if-else in golang)
	year := 2004
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				fmt.Println(year, " - It is a leap year")
			} else {
				fmt.Println(year, " - It is not a leap year")
			}

		} else {
			fmt.Println(year, " - It is a leap year")
		}
	} else {
		fmt.Println(year, " - It is not a leap year")
	}

	//single if condition leap year
	if year%100 != 0 && year%4 == 0 || year%400 == 0 {
		fmt.Println(year, " - It is a leap year")
	} else {
		fmt.Println(year, " - It is not a leap year")
	}
}

func bubbleSort() {
	//bubble sorting
	arr := []int{20, 10, 40, 50, 30, 70, 60}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j] //similar to python swapping
			}
		}
	}
	fmt.Println("bubble Sort", arr)
}

// fibonacci series using recursion
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func removeDuplicatesSlice() {
	//removes duplicates using empty array and python like "in operator" function
	num_arr := []int{23, 45, 21, 45, 21, 45, 67, 34, 2, 21, 2, 0, 45}
	duplicate_removal := []int{}
	for i := range num_arr {
		if !(sliceContains(duplicate_removal, num_arr[i])) {
			duplicate_removal = append(duplicate_removal, num_arr[i])
		}
	}

	fmt.Println(duplicate_removal)
}

func sliceContains(arr []int, value int) bool { // python's "in" "not in" functionality
	for i := range arr {
		if arr[i] == value {
			return true
		}
	}
	return false
}

func removeSliceValuesAnotherMethod() {
	//"different method, remove element by index but slice order changed"
	a1 := []string{"A", "B", "C", "D", "E"}
	remove_index := 2
	a1[remove_index] = a1[len(a1)-1]                          //{"A", "B", "E", "D", "E"} copy last value into removing index value
	a1 = a1[:len(a1)-1]                                       //After already copied last value into remove_index, now removes the duplicate last value by slicing
	fmt.Println("remove element but slice order changed", a1) //[A B E D]
}

func removeSliceValues() {
	//removes one or multiple occurence values in slice
	u := []int{3, 5, 3, 6, 7, 3}
	f := len(u)
	for i := 0; i < f; i++ {
		fmt.Println("how many time loop runned despite removing values")
		if u[i] == 3 { //removes 3
			u = append(u[:i], u[i+1:]...) //if we have to remove one occurence value, this line code is enough
			f = f - 1                     //whenever we removes the value in runtime, should reduce the slice length by 1
		}
	}
	fmt.Println("after removed one or multiple occurence value in slice", u)
}
func IsSlicesorted() {
	// program to check whether slice is already sorted
	f1 := []int{1, 2, 3, 4, 10, 101}
	for i := 0; i < len(f1)-1; i++ {
		if f1[i] > f1[i+1] {
			fmt.Println("slice is not sorted")
			break
		}
	}
	fmt.Println("slice is sorted")

}

func checkRomanLetter() {
	number := 1988
	finaloutput := ""
	tmpoutput := ""

	for number > 0 {
		tmpoutput, number = validateRomanLetter(number)
		finaloutput = finaloutput + tmpoutput
	}

	fmt.Println("Roman letter for given input", finaloutput)
}

func validateRomanLetter(number int) (string, int) {

	remainder := -1
	quotient := -1
	out := ""

	for _, j := range romanLettersInDescending {
		quotient = number / j
		remainder = number % j
		if quotient > 0 {
			for quotient > 0 {
				out = out + romanMap[j]
				fmt.Println("---", quotient)
				quotient--
			}
			return out, remainder
		}
	}

	return "", 0
}

func ValidateEventInt64Type(value interface{}) int64 {

	switch reflect.ValueOf(value).Kind() {
	case reflect.String:
		EventInt64Value, _ := strconv.ParseInt(value.(string), 10, 64)
		return EventInt64Value
	case reflect.Int:
		return int64(value.(int))
	case reflect.Int64:
		return int64(value.(int64))
	case reflect.Float32:
		return int64(value.(float32))
	case reflect.Float64:
		return int64(value.(float64))
	}

	return 0

}

// returns maximum value from slice
func maximumInt(v []int) int {
	max := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > max {
			max = v[i]
		}
	}

	return max
}

func printnumbers(n int) {
	if n > 0 {
		printnumbers(n - 1)
		fmt.Println(n)

	}
}

// count all substring occurances
func findSubStringAllOccurences(a, b string) int {
	substringLength := len(b)
	j := 0
	totalOccurences := 0
	for i := 0; i < len(a); i++ {
		j = 0
		if a[i] == b[j] {
			//this while loop for runs for length substring

			for j < substringLength {
				if a[i] != b[j] {
					break
				}
				j++
				//ensures array not going beyond the string length
				if i >= len(a) {
					break
				}
				i++
			}

			fmt.Println("substring found indexes in string", i-substringLength)
			totalOccurences++
		}
	}

	return totalOccurences
}

func findSubString(a, b string) bool {
	substringLength := len(b)
	j := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[j] {
			//this while loop for runs for length substring
			for j < substringLength {
				if a[i] != b[j] {
					return false
				}
				j++
				//ensures array not going beyond the string length
				if i >= len(a) {
					break
				}
				i++
			}

			return true //substring is there and code hits here means it true case.
		}
	}

	return false
}

func maximumNumbers(n []int) (int, int) {
	maxCount := 0
	count := 0
	index := -1
	for i := 0; i < len(n); i++ {
		count = 1
		for j := i + 1; j < len(n); j++ {
			if n[i] == n[j] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			index = i
		}
	}

	return n[index], maxCount

}

func sortInterfaceSlice() {
	s := []interface{}{"mat", "bat", "rat", "pat", 6, 3, 1, -1, 4, 5}

	stringSlice := []string{}
	intSlice := []int{}
	for _, i := range s {
		switch i.(type) {
		case int:
			intSlice = append(intSlice, i.(int))
		case string:
			stringSlice = append(stringSlice, i.(string))
		}
	}

	sortingWithGenerics(intSlice)
	sortingWithGenerics(stringSlice)

	for i := 0; i < len(intSlice); i++ {
		s[i] = intSlice[i]
	}

	for j := 0; j < len(stringSlice); j++ {
		s[len(intSlice)+j] = stringSlice[j]
	}

	fmt.Println("sorted the input interface slice output -- ", s)
}

func sortingWithGenerics[slice int | string](s []slice) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

func SortedMaximumNumbers() {

	df := []int{1, 2, 2, 2, 2, 3, 3, 5, 5, 5, 5, 5, 5}
	j := 0
	count := 0
	max := 0
	index := -1
	for i := 1; i < len(df); i++ {
		if df[i] == df[j] {
			count++
		} else {
			if count > max {
				max = count
				index = j
				count = 0
			}

		}
		j = j + 1
	}

	//handles the last matched values
	if count > max {
		max = count
		index = j
	}

	fmt.Println("maximum times of number in sorted slice ", df[index], max)

}

func removeDuplicatesInSortedString() {
	df := "sathissssaaaaaasssssssshttttttttt"
	s := strings.Split(df, "")
	sort.Strings(s)
	o := "" + s[0] //we needs to add the missing first charcter, because we adding only s[i+1]
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			o = o + s[i+1]
		}
	}

	fmt.Println(s, o)
}

func NumbersInExpectedSum() {
	// df := []int{1, 2, 3, 0, 3}
	df := []int{0, 1, -1, 0}
	out := [][]int{}
	ExpectedSum := 3
	for i := 0; i < len(df); i++ {
		for j := i + 1; j < len(df); j++ {
			if df[i]+df[j] == ExpectedSum { //if equal to the expectedsum
				tmp := []int{}
				tmp = append(tmp, df[i])
				tmp = append(tmp, df[j])
				out = append(out, tmp)
			}
		}
	}

	fmt.Println("output----", out, len(out)) //output---- [[1 2] [3 0] [0 3]]
}

func largestPossibleNumber(arr1 []int) string {

	if len(arr1) == 0 {
		return "Empty Slice"
	}

	arr := []string{}

	for _, j := range arr1 {
		arr = append(arr, strconv.Itoa(j))
	}

	if len(arr) == 1 {
		return arr[0]
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			//concats the slice string values and compares it
			if arr[i]+arr[j] < arr[j]+arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return strings.Join(arr, "")

}

func missingNumbersInSortedSlice(d []int) {
	for i := 0; i < len(d)-1; i++ {
		tmp := 0
		if d[i+1]-d[i] > 1 {
			tmp = d[i] + 1
			//while loop, until we prints the missing values between two values.
			for tmp < d[i+1] {
				fmt.Println(tmp)
				tmp = tmp + 1
			}
		}
	}
}

// Recursion examples
func power(value int, pow int) int {
	if pow == 1 {
		return value
	}
	return value * power(value, pow-1)
}

func factorial(n int) int {
	fmt.Println(n)
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func fiaboncii(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	return fiaboncii(n-1) + fiaboncii(n-2)
}

// ReverseString recursion examples
func ReverseString(s string) string {
	j := 0
	i := len(s) - j
	c := ""
	if i == 0 {
		return s
	} else {
		c = "" + ReverseString(string(s[i]))
		j = j - 1
		return c
	}
}

func reverseString(s string, substring string, length int) string {
	if length == 0 {
		return string(s[0])
	}

	return substring + reverseString(s, s[length-1:length], length-1)
}

func fibonacci() {

	aa := 0
	bb := 1
	cc := 0
	fmt.Printf("fiaboncii-\n%d\n%d", aa, bb)
	for i := 0; i < 10; i++ {
		cc = aa + bb
		fmt.Print("\n", cc)
		aa = bb
		bb = cc
	}

}

func reverseStringWithoutAnotherString() {
	ee := "sathish"
	numSlice := []rune(ee) //rune -- converts string into slice of ascii character
	for i := 0; i < len(numSlice)/2; i++ {
		j := len(numSlice) - 1 - i
		numSlice[j], numSlice[i] = numSlice[i], numSlice[j]
	}
	//finally converts back rune into string
	fmt.Println(string(numSlice))
}

func reverseSliceWithoutAnotherSlice() {
	numSlice := []int{11, 12, 13, 14, 15, 16, 17, 18}
	for i := 0; i < len(numSlice)/2; i++ {
		j := len(numSlice) - 1 - i
		// numSlice[i], numSlice[j] = numSlice[j], numSlice[i]
		numSlice[j], numSlice[i] = numSlice[i], numSlice[j]
	}
	fmt.Println(numSlice)
}

func InsertIndexSlice() {
	numSlice := []int{11, 12, 13, 14, 15}
	Insertindex := 2
	InsertValue := 99
	numSlice = append(numSlice[:Insertindex+1], numSlice[Insertindex:]...)
	numSlice[Insertindex] = InsertValue
	fmt.Println(numSlice)
}

func removeValueSlice() {
	numSlice := []int{11, 12, 13, 14, 15}
	index := 2
	numSlice = append(numSlice[:index], numSlice[index+1:]...)
	fmt.Println(numSlice)
}

func mapOfSlice() {
	mapSlice := map[string][]string{}
	mapSlice["1"] = []string{"a", "aa"}
	mapSlice["2"] = []string{"b", "bb"}
	mapSlice["2"] = append(mapSlice["2"], "bbb")
	mapSlice["2"] = append(mapSlice["2"], mapSlice["1"]...)
	for key, value := range mapSlice {
		fmt.Println(key, value)
	}
}

// copying oldkey with newkay and then deletes old key
func renameMapkeyName() {
	dd := map[string]string{"x": "1", "y": "2"}
	oldKey := "x"
	newKey := "a"
	dd[newKey] = dd[oldKey]
	fmt.Println("renaming map key", dd)
	delete(dd, "x")
	fmt.Println("renaming map key", dd)

}

// two D map
func twoDimensionMap() {
	twoDMap := map[string]map[string]string{}
	twoDMap["a"] = map[string]string{} //we need to initialize empty map like this
	twoDMap["b"] = map[string]string{}
	twoDMap["a"]["1"] = "ssss"
	twoDMap["b"]["2"] = "dddd"
	fmt.Println(twoDMap)
}

func intstringslicesumOfDigits() {
	dr := 5678
	str := strconv.Itoa(dr)
	strSlice := strings.Split(str, "")
	sum := 0
	for i := 0; i < len(strSlice); i++ {
		p, _ := strconv.Atoi(strSlice[i])
		sum = sum + p
		fmt.Println(strSlice[i], sum, reflect.TypeOf(strSlice[i]))
	}

}

func perfectNumber(number int) {
	sum := 0
	for i := 1; i < number; i++ {
		if number%i == 0 {
			sum = sum + i
		}
	}

	if sum == number {
		fmt.Println("it is a perfect number", number)
	} else {
		fmt.Println("Not perfect number", number)
	}
}

func primeNumbers(number int) {
	flag := false
	for i := 2; i < number+1; i++ {
		flag = false
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("prime numbers------", i)
		}
	}
}

// easy method
func reverseStringV1() {
	ty := "sathish have a nice day"
	final := ""
	for i := len(ty) - 1; i >= 0; i-- {
		final = final + string(ty[i])
	}

	fmt.Println("reverse----", final)
}

func countWords() {
	er := map[string]int{}
	testString := "sathish have a nice day"
	fmt.Println("pppp", len(testString))
	for _, j := range testString {
		if string(j) == " " {
			continue
		}
		//This also works
		// if er[string(j)] > 0 {
		// 	er[string(j)] = er[string(j)] + 1
		// } else {
		// 	er[string(j)] = 1
		// }

		//This is also works, But this is used wodely
		if _, k := er[string(j)]; k {
			er[string(j)] = er[string(j)] + 1
		} else {
			er[string(j)] = 1
		}
	}

	fmt.Println("countwords in map", er)

}

func sum_of_digits(number int) {
	f1 := 0
	for number > 0 {
		temp := number % 10
		f1 = f1 + temp
		number = number / 10

	}

	fmt.Println("sum of all digits", f1)
}

func reverseNumber(number int) {
	reverse := 0
	for number > 0 {
		remainder := number % 10
		// f1 = f1 + temp
		reverse = reverse*10 + remainder
		number = number / 10
		fmt.Println(number)

	}

	fmt.Println("reverse number", reverse)
}

func number_of_digits(number int) {
	c := 0

	// count the int number digits
	for number > 0 {
		number = number / 10
		c = c + 1

	}

	fmt.Println("no of digits", c)
}

func ReadWriteFile() {
	filename := "ccc.txt"
	filedata := []byte("good hello this is text data added into the file new version data Go has excellent built-in support for file operations. Using the os package, you can easily open, read from, write to and close the file.In this example, we focus on writing data to a file. We show you how you can write text and binary data in different ways - entire data at once, line by line, as an array of bytes in a specific place, or in a buffered manner")
	err := os.WriteFile(filename, filedata, 0666) //clears existing file data and adds the data
	fmt.Println(err)

	readfile, err := os.ReadFile(filename)
	fmt.Println("read file data from OS package----------------", err, string(readfile))

	//By using this function, We can set different flags like create file, readonly,writeonly,readAndWrite,Append
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //os.O_CREATE - if file not there it creates new one
	file.WriteString("adding new data =======================================")
	readfile, err = os.ReadFile(filename)
	fmt.Println("appended new data into file from OS package----------------", err, string(readfile))

	//example for reading file
	file, err = os.OpenFile(filename, os.O_RDONLY, 0666)
	fmt.Println("Againing reading the data----------------", err, string(readfile))

	var lines = []string{
		"Go",
		"is",
		"the",
		"best",
		"programming",
		"language",
		"in",
		"the",
		"world",
	}

	//bufio : Buffered Input/Output --
	//bufio I/O -- It accumulates data upto the buffer size and then doing the flush(for performs acutal I/O operation) once. So it reduces the number of system calls and works efficiently overall.

	//Example bufio buffer write operation-- It accumulate(collects without directly write data on file), So it reduces the small data write system calls into files and works efficiently overall.

	//default buffer size is 4096K, Bufio can hold the data upto this capacity

	//bufio write
	f, err := os.Create("bufiofile.txt")
	bufioWriter := bufio.NewWriter(f) //have to use the created file from "os" package, This also clears existing file data

	i, err := bufioWriter.Write(filedata)
	fmt.Println(i, err)
	for _, v := range lines {
		//can do each write as different data type
		bufioWriter.Write([]byte(v))  //write data as byte data type
		bufioWriter.WriteString("\n") //Adds new line //write data as string data type
	}
	err1 := bufioWriter.Flush() //After doing flush(), Then only it writes data on the file until buffer default capacity size 4096K.
	fmt.Println(err1)
	f.Close() //should close it, If we are going to read it immeditely

	//reads bufio
	readfile1, _ := os.Open("bufiofile.txt")
	bufioScanner := bufio.NewScanner(readfile1)

	for bufioScanner.Scan() {
		fmt.Println("read file data----------", bufioScanner.Text()) //reading file data line by line
	}

}

func SampleRetryLogic(attempts int, f func() error) (err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			fmt.Println("retrying after error:", err)
			time.Sleep(2 * time.Second)
		}
		err = f() //Calling the callback function with in retry attempt count.
		if err == nil {
			return nil //If no error, Returns from function
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func SortMapValuesUsingSeperateKeyValueSlices() {
	m := map[int]int{10: 00, 6: 66, 1: 11, 2: 22, 3: 33, 4: 44, 0: 55}
	keySlice := make([]int, len(m))
	valueSlice := make([]int, len(m))

	p := 0
	for k, v := range m {
		keySlice[p] = k
		valueSlice[p] = v
		p++
	}

	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(m); j++ {
			if valueSlice[i] < valueSlice[j] { //Reverse sort the map values through slice
				valueSlice[i], valueSlice[j] = valueSlice[j], valueSlice[i]
				keySlice[i], keySlice[j] = keySlice[j], keySlice[i] //Swaps map keys as well while swaps the map values, So we maintains the same order of both map keys and values on keySlice and ValueSlice
			}
		}
	}

	fmt.Println(keySlice, valueSlice)

	for _, i := range keySlice { //Iterating the map in sorted order
		fmt.Println("sorted in map values with key", i, m[i])
	}

}

func StackImplmentation() {

	s := Stack{}
	e := s.Pop()
	if e != nil { //Here checking nil return value
		fmt.Println(*s.Pop())
	}
	s.Push(10)
	s.Push(20)
	fmt.Println(*s.Pop())
	s.Push(100)
	s.Push(200)
	fmt.Println(*s.Pop())
	fmt.Println(s)
}

type Stack struct {
	array []int
}

func (s *Stack) Push(value int) {
	s.array = append(s.array, value)
}

func (s *Stack) Pop() *int {
	if len(s.array) == 0 {
		return nil
	}
	last := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return &last
}

func RemoveOneorMultipleSubstringUsingSlice() {
	s := "hisTutorials his Freak his hello his"
	removeSubString := "his"
	flag := false
	for i := 0; i < len(s); i++ {
		if string(s[i]) == string(removeSubString[0]) {
			flag = false
			for j := 0; j < len(removeSubString); j++ {
				if s[i+j] != removeSubString[j] { //Compares each character in both string and substring
					flag = true
					break
				}
			}
			if !flag {
				tmp := i + len(removeSubString)
				s = s[:i] + s[tmp:] //We found the substring match, So removing single substring occurance using slicing
				i = tmp             //increasing the i value for substring length for unnneeded iteration processing.
			}

		}
	}

	fmt.Println(len(s), s)

}

func RemoveSpaceUsingStringSlicing() {
	s := "Tutor ials Freak hi hello"
	fmt.Printf("before memory address - %p", &s)
	fmt.Println()
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			s = s[:i] + s[i+1:]
		}
	}

	fmt.Println(len(s), s)
	fmt.Println()
	fmt.Printf("After memory address - %p", &s)
}

func RemoveSpaceUsingSlice() {
	s := "Tutor ials Freak hi hello"
	strSlice := make([]string, len(s))

	j := 0
	for _, i := range s {
		if string(i) != " " {
			strSlice[j] = string(i)
			j++
		}
	}
	s1 := strings.Join(strSlice, "") //strings join() automatically removes the last empty length in slice
	fmt.Println(len(strSlice), strSlice, len(s1), s1)
	fmt.Println(strings.Count(s, " "))

}

func sumOfMissingNumbers() {
	array := []int{4, 3, 8, 1, 2, 10, 12}
	sort.Ints(array) //sort the slice first
	sum := 0
	fmt.Println(array)
	for i := 0; i < len(array)-1; i++ {
		tmp := array[i+1] - array[i]

		if tmp > 1 { //If the difference is greater than 1, Then we have the missing numbers,
			fmt.Println(tmp, sum, array[i]+1)
			for j := array[i] + 1; j < array[i]+tmp; j++ {
				sum = sum + j
			}
		}
	}

	fmt.Println("Sum of Missing number", sum)
}

func SortStringsinSlice() {
	s := []string{"watermelon", "banana", "apple", "cherry", "date", "orange", "elderberry", "mango", "grape"}

	//follows same bubble sort logic
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if string(s[i][0]) > string(s[j][0]) { //s[j][0]  - [0] constantly compares first letter in all strings
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	fmt.Println("sorted with first letter in slice values", s) //[apple banana cherry date elderberry grape mango orange watermelon]
}

type Student struct {
	name  string
	score int
}

func SortByStructfields() {
	students := []Student{

		Student{name: "John", score: 87},
		Student{name: "Albert", score: 68},
		Student{name: "Bill", score: 68},
		Student{name: "Sam", score: 98},
		Student{name: "Xenia", score: 87},
		Student{name: "Lucia", score: 87},
		Student{name: "Tom", score: 91},
		Student{name: "Jane", score: 68},
		Student{name: "Martin", score: 71},
		Student{name: "Julia", score: 87},
	}

	sort.SliceStable(students, func(i, j int) bool {
		return students[i].score > students[j].score //students score in descending order.
	})

	fmt.Println("sorted struct by fields ", students)

}

func isValid(str string) bool {
	var s []byte

	if str[0] == '(' || str[0] == '{' || str[0] == '[' {

		for i := 0; i < len(str); i++ {
			s = append(s, str[i])
			if len(s) > 1 {
				fmt.Println(s, s[len(s)-2])
				if s[len(s)-2] == '[' && s[len(s)-1] == ']' {
					s = s[:len(s)-2]
				} else if s[len(s)-2] == '{' && s[len(s)-1] == '}' {
					s = s[:len(s)-2]
				} else if s[len(s)-2] == '(' && s[len(s)-1] == ')' {
					s = s[:len(s)-2]
				}
			}
		}
		if len(s) == 0 {
			return true
		}
	}
	return false
}

func CheckAnagram() {

	string1 := "a gentleman"
	string2 := "elegant man"
	t := strings.Split(string1, "")
	t1 := strings.Split(string2, "")

	sort.SliceStable(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	sort.SliceStable(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})

	// sort.Strings(t)
	// sort.Strings(t1)

	if strings.Join(t, "") == strings.Join(t1, "") {
		fmt.Println("Both strings are anagram")
	} else {
		fmt.Println("Both strings are not anagram")
	}
}

func CheckAnagramV1() {
	string1 := "a gentleman"
	string2 := "elegant man"
	t := []byte(string1)
	t1 := []byte(string2)

	sort.SliceStable(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	sort.SliceStable(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})

	if string(t1) == string(t) {
		fmt.Println("Both strings are anagram")
	} else {
		fmt.Println("Both strings are not anagram")
	}
}

func CheckAnagramUsingMaps() {
	string1 := "a gentleman"
	string2 := "elegant man"
	map1 := map[string]int{}
	map2 := map[string]int{}

	for i := 0; i < len(string1); i++ {
		if _, found := map1[string(string1[i])]; !found {
			map1[string(string1[i])] = 1
		} else {
			map1[string(string1[i])] = map1[string(string1[i])] + 1
		}
		if _, found := map2[string(string2[i])]; !found {
			map2[string(string2[i])] = 1
		} else {
			map2[string(string2[i])] = map2[string(string2[i])] + 1
		}
	}

	fmt.Println(map1, map2, len(map1), len(map2))

	fmt.Println("comparing two maps---", reflect.DeepEqual(map1, map2))

	flag := false
	for key, value := range map1 {
		maps2Value, found := map2[key]
		if maps2Value != value || !found {
			fmt.Println("Not anagram found in both strings")
			flag = true
			break
		}
	}

	if !flag {
		fmt.Println("anagram found in both strings")
	}

}

func RemovecharactersToMakeAnagram() {
	str1 := "bcadeh"
	str2 := "hea"
	biggerStr := str1
	lStr := str2
	if len(str1) < len(str2) {
		biggerStr = str2
		lStr = str1
	}

	for _, i := range biggerStr {
		if !strings.Contains(lStr, string(i)) {
			fmt.Println("This character needs to remove in biggerString", string(i))
		}
	}

}

// Can use the same logic for second minimum numbers(just have to change the logic signs <>) in slice
func secondMaximumWithoutSorting() {
	s1 := []int{23, 45, 2, 56, 51, 10, 15, 50}
	max1 := 23
	max2 := 45
	if max1 < max2 {
		max1, max2 = max2, max1
	}

	for i := 2; i < len(s1); i++ {
		if s1[i] < max2 {
			continue
		}

		//have value greater than max1, So assigns max1 value to max2, s[i] to max1.
		if s1[i] > max1 {
			max2 = max1
			max1 = s1[i]
			continue
		}

		if s1[i] > max2 && s1[i] < max1 {
			max2 = s1[i]
		}

	}

	fmt.Println("Second largest in slice", max2)
}

func BinarySearch() {
	array := []int{1, 3, 4, 5, 8, 9, 10, 13}
	SearchingValue := 4
	high := len(array) - 1
	low := 0

	for low <= high {
		mid := high - low/2
		if array[mid] == SearchingValue { //Keep on adjusting the low and high values by comparing the mid value and then matches values in this condition.
			fmt.Println("value SearchingValue at index", mid)
		}
		if array[mid] < SearchingValue {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

}

func commonUncommonSliceValues() {
	df := []int{1, 2, 3, 4, 5, 6}
	fg := []int{2, 3, 4, 7, 8, 10, 5}
	CommonValues := []int{}
	unCommonValues := []int{}
	// flag := false

	for i := 0; i < len(df); i++ {
		flag := false
		for j := 0; j < len(fg); j++ { //should start this for loop with 0th index everytime, Because we may have two input slices with different lenghts

			//condition passes for common values between two slices
			if df[i] == fg[j] {
				CommonValues = append(CommonValues, df[i])
				flag = true
			}

		}

		if !flag {
			unCommonValues = append(unCommonValues, df[i])
		}
	}

	fmt.Println("common values in two slices", CommonValues)      //[2 3 4 5]
	fmt.Println("un common values in two slices", unCommonValues) //[1 6] this values not present in second input slice -- {2, 3, 4, 7, 8, 10, 5}

}

func MatrixPrograms() {
	MatrixArray := [3][3]int{} //2D arrays
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			MatrixArray[i][j] = j + 10 //assigns value
		}
	}

	fmt.Println("Matrix using array", MatrixArray)

	TwoMatrix := make([][]int, 5) //Always use make() for creating 2D slice, Normal empty slice creation EX:[][]int{} causes problems
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			TwoMatrix[i] = append(TwoMatrix[i], j) //appends matrix value
		}
	}

	fmt.Println("Matrix using Slice", TwoMatrix)

	//copy uneven matrix to another matrix
	df := [][]int{{2, 3, 4}, {5, 6}, {7, 8, 9, 0}}
	fmt.Println(len(df), len(df[0]), len(df[2]))
	copy_Matrix := make([][]int, len(df))
	for i := 0; i < len(df); i++ {
		//len(df[i]) -- range depends upon particular row length
		for j := 0; j < len(df[i]); j++ {
			copy_Matrix[i] = append(copy_Matrix[i], df[i][j])
		}
	}

	fmt.Println("copied matrix", copy_Matrix)
	copy_Matrix = [][]int{}
	fmt.Println("empting matrix", copy_Matrix, len(copy_Matrix))

	df1 := [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 0}}
	for i := 0; i < len(df1); i++ {
		for j := 0; j < len(df1[i]); j++ {
			fmt.Println("default row_iteration", df1[i][j])
		}
	}

	for i := 0; i < len(df1); i++ {
		for j := 0; j < len(df1[i]); j++ {

			//df1[j][i]
			//-- 0,0 - 1,0 - 2,0
			//-- 0,1 - 1,1 - 2,1
			//-- 0,2 - 1,2 - 2,2
			fmt.Println("Matrix Column_iteration", df1[j][i])
		}
	}

	//transpose matrix --
	//Rows values of one matrix should be equal to the columns values of another matrix (or) Changing rows values of matrix to columns values of matrix, We can use the same above rows-column iteration logic
	//Ex: df1[i][j] == df1[j][i]

	//sort matrix:
	//Have temp 1D slice
	//iterate input matrix, Append all values into temp
	//sort temp slice
	//iterates and Assigns the sorted temp slice values back into input matrix

	om := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 0, 1, 2},
		{3, 4, 5, 6}}

	rotateMatrixBy90Degrees := [4][4]int{}
	k := 0
	for i := len(om) - 1; i >= 0; i-- {

		for j := 0; j < len(om); j++ {
			rotateMatrixBy90Degrees[j][i] = om[k][j] //Here we are not doing simple appending, So we can't create otateMatrixBy90Degrees with make() and assign the values directly, First have to create the same size matrix with dummy values, Then we can assign like this. To avoiding this now,"rotateMatrixBy90Degrees" created with fixed size array.

		}
		k++
	}

	fmt.Println("rotateMatrixBy90Degrees", rotateMatrixBy90Degrees) //       rotateMatrixBy90Degrees
	// [[3 9 5 1]
	// [4 0 6 2]
	// [5 1 7 3]
	// [6 2 8 4]]
}

func NonConsectiveCharacters() {
	s := "sdddffffghhhklp"
	tmp := ""
	for i := 0; i < len(s)-1; i++ {
		tmp = string(s[i+1])
		if string(s[i]) != string(tmp) {
			fmt.Print(string(s[i]))
		}
	}
	fmt.Print(string(tmp)) //This prints last character
	//final output -- "sdfghklp"
}

// fmt.Println(concat(1, 2, 3, "a", "b", "c", []int{22, 34}, 23.45, []string{"dd", "dg"}))
func concat(input ...interface{}) string {
	var output string
	for _, i := range input {
		switch i.(type) {
		case string:
			output = output + i.(string)
		case int:
			output = output + fmt.Sprintf("%v", i.(int))
		default:
		}
	}

	return output //123abc
}

func fibonacciSeriesInputValueShortestBetweenRange() {
	//0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181  - fibonacci Series upto 5000

	input := 350 //If input is equal to fibonacci series number, That should be printed. If input is between two fibonacci series number EX: 350 is between 377 and 610. So the nearest value between two values should be printed.
	a := 0
	b := 1
	c := 0
	for i := 0; i < 20; i++ {
		c = a + b

		if input == c {
			fmt.Println(c)
			break
		}

		if input > b && input < c {
			tmp := input - b
			tmp1 := c - input

			if tmp > tmp1 {
				fmt.Println(c)
			} else {
				fmt.Println(b)
			}
			break
		}
		a = b
		b = c

	}
}

func CheckStringPrefixwithSlice() {
	s1 := []string{"i", "love", "leetcode", "apples"}

	s2 := "iloveleetcode"
	// s2 := "i"
	flag := 0
	for i := 0; i < len(s1); i++ {
		tmp := s1[:i]
		if strings.Join(tmp, "") == s2 {
			fmt.Println(strings.Join(tmp, ""), s2)
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Println("prefix found")
	} else {
		fmt.Println("prefix not found")
	}
}

func CheckEveryAlphabetCharactersFoundinInput() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	fmt.Println()
	flag := 0
	for i := 97; i < 123; i++ {
		flag = 0
		for j := 0; j < len(sentence); j++ {
			if string(i) == string(sentence[j]) {
				flag = 1
				break
			}
		}

		if flag == 0 {
			fmt.Println("Alphabet Character is not found in the input ---", string(i))
			break
		}

	}

	if flag == 1 {
		fmt.Println("All Alphabet Character are found in the input")
	}
}

func SummationOfTwoStringsWithThirdString() {

	firstWord := "acb"
	secondWord := "cba"
	targetWord := "cdb"

	// firstWord := "aaa"
	// secondWord := "a"
	// targetWord := "aaaa"

	sum := 0
	sum1 := 0
	sum2 := 0
	for i := 0; i < len(firstWord); i++ {
		tmp := int(firstWord[i]) - 97
		sum = sum + tmp
	}

	for i := 0; i < len(secondWord); i++ {
		tmp := int(secondWord[i]) - 97
		sum1 = sum1 + tmp
	}

	for i := 0; i < len(targetWord); i++ {
		tmp := int(targetWord[i]) - 97
		sum2 = sum2 + tmp
	}

	fmt.Println(sum, sum1, sum2, "-------------", sum+sum1 == sum2)
}

func compareMaps() {

	input := map[string]int{"a": 4, "b": 5, "c": 6}
	output := map[string]int{"a": 4, "b": 5, "c": 6}
	if len(input) != len(output) {
		panic("not equal")
	}
	for key, value := range input {
		if _, ok := output[key]; !ok {
			panic("Key is not found, not equal")
		}
		if value != output[key] {
			panic("Key is found, but value not equal")
		}
	}

	fmt.Println("Both maps are equal")
}

// not working need to check for some edge cases
func IndexStartSubString() {
	haystack := "hello"
	needle := "lo"

	index := -1
	tmp := -1
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			index = i
			tmp = i
			for k := 0; k < len(needle); k++ {
				if i+k >= len(haystack) {
					break
				}
				if haystack[i] != needle[k] {
					fmt.Println("not matched")
					index = -1
					break
				}
				tmp++
			}

		}
	}

	fmt.Println(index)
}

func SmallestMissingNumberWithoutSorting() {
	// s := []int{3, 4, -1, 2, 1}
	s := []int{1, 2, 0}
	Maximum := slices.Max(s)
	for i := 1; i <= Maximum+1; i++ { //Gives maximum with 1 one more value range, If no missing numbers from 1 to actual maximum value in the slice
		if !slices.Contains(s, i) {
			fmt.Println("Smallest missing number", i)
			break
		}
	}

}

func FindWhichDay() {
	days := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	Fromday := "wed" //Which day would be, Assume today is Wednesday, So 10 days from today's wednesday
	numOfDays := 10
	// numOfDays := 23

	StartingDay := slices.Index(days, Fromday)
	startingPosition := StartingDay + 1 //Getting starting day's index

	if numOfDays > len(days) {
		numOfDays = numOfDays % len(days) //Here we gets remainder, So the remainder used to get the next days.
	}

	index := -1
	//circular iteration starts from startingPosition index
	for i := 0; i < numOfDays; i++ {
		index = (i + startingPosition) % len(days)
	}

	fmt.Println(days[index]) //here using the last index value from baove for loop last iteration

}

func TwoDimensionalSliceFindOverlapRangeValues() {
	s := [][]int{{1, 5}, {3, 4}, {6, 9}, {7, 8}, {2, 3, 10}, {0}, {0, 1}, {3, 9}, {6, 20}, {15, 16}}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i][0] < s[j][0] && s[i][len(s[i])-1] > s[j][len(s[j])-1] {
				fmt.Println(s[i], s[j])
				// [1 5] [3 4]    //{3,4} -- Values range located between [1, 5] number range
				// [6 9] [7 8]
				// [2 3 10] [3 9]
				// [6 20] [15 16]
			}
		}

	}
}

func RemoveDuplicateNumbersIn2DSlice() {

	s := [][]int{{1, 5}, {3, 4}, {6, 7}, {8, 9}, {2, 3, 5}, {0}, {0, 1}, {1, 11}, {6, 20}, {15, 16}}
	output := [][]int{}

	output = append(output, s[0]) //loads the first value for finding duplicates from index 1
	for i := 1; i < len(s); i++ {

		if !TwoDContains(output, s[i]) {
			output = append(output, s[i])
		}
	}

	fmt.Println(output) //[[1 5] [3 4] [6 7] [8 9] [0] [15 16]]

}

func TwoDContains(output [][]int, s []int) bool {
	for i := 0; i < len(s); i++ {
		for k := 0; k < len(output); k++ {
			for z := 0; z < len(output[k]); z++ {
				if output[k][z] == s[i] {
					return true
				}
			}
		}
	}

	return false
}
func Sort2DSliceByfirstValue() {
	s := [][]int{{1, 5}, {3, 4}, {6, 7}, {8, 9}, {2, 3, 5}, {0}, {0, 1}}
	k := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i][k] > s[j][k] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	fmt.Println("ascending order", s) //[[0] [0 1] [1 5] [2 3 5] [3 4] [6 7] [8 9]]

	//2D sort using sliceStable
	sort.SliceStable(s, func(ii, jj int) bool {
		return s[ii][k] > s[jj][k]
	})

	fmt.Println("descending order", s) //[[8 9] [6 7] [3 4] [2 3 5] [1 5] [0] [0 1]]
}

func Sort2DSliceBysumOfValues() {
	s := [][]int{{1, 5}, {3, 4}, {6, 7}, {8, 9}, {2, 3, 5}, {0}, {0, 1}}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if SumOfSlice(s[i]) > SumOfSlice(s[j]) {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	fmt.Println(s)
}

func SumOfSlice(s []int) int {
	sum := 0
	for _, i := range s {
		sum = sum + i
	}
	return sum
}

func CombinationValuesWithGivenTarget() {
	input1 := []int{10, 1, 2, 7, 6, 1, 5}
	target := 10
	output := [][]int{}
	for i := 0; i < len(input1); i++ {
		if input1[i] == target {
			output = append(output, []int{input1[i]})
			continue
		}

		sum := 0
		tmp := []int{}
		for j := i + 1; j < len(input1); j++ {
			fmt.Println(sum, tmp)
			if sum+input1[j] == target {
				tmp = append(tmp, input1[j])
				output = append(output, tmp)
				break
			}
			if sum+input1[j] > target {
				continue
			}
			sum = sum + input1[j]
			tmp = append(tmp, input1[j])

		}
	}

	fmt.Println(output) //[[10] [2 7 1]]  -- Combination of values of given target
}

func SumOfTripletsLessThanTarget() {
	num := []int{-2, 0, 1, 3, 4, 5, 6, 7}
	// output := make([][]int, len(num))
	output := [][]int{}
	target := 10
	for i := 0; i < len(num)-2; i++ {
		sum := num[i] + num[i+1] + num[i+2]
		if sum < target {
			output = append(output, []int{num[i], num[i+1], num[i+2]})
			// output[i] = []int{num[i], num[i+1], num[i+2]}
			fmt.Println(num[i], num[i+1], num[i+2])
		}
	}

	fmt.Println("sum of triplets less than target", output)

}

func SumOfMissingConsectiveNumbers() {
	in := []int{1, 2, 3, 7, 8, 13, 14, 16}
	sum := 0
	for i := 0; i < len(in)-1; i++ {
		if in[i+1]-in[i] > 1 {
			tmp := in[i] + 1
			for tmp < in[i+1] {
				sum = sum + tmp
				fmt.Println(tmp, sum)
				tmp++
			}
		}
	}

	fmt.Println(sum)
}
