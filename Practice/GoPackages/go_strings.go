package main

import (
	"fmt"
	"reflect"
	s "strings" //import "strings" as "s"
	"unicode"
)

var p = fmt.Println //using "fmt.Println" as "p" for convenience

func main() {

	// xy1 := "hello"
	// xy1[2] = "k"      //like python, We can't modify the strings directly in golang
	// fmt.Println(xy1)

	// multi -line string declaration
	str := `This is a
multiline
string.`

	fmt.Println(str)

	p("abc" + "edf")

	//strings module inbuilts
	p("Contains:  ", s.Contains("test", "es")) // python's "in"
	p(s.ContainsRune("test", 't'))             //check rune substring contains in string or not

	p("Count:     ", s.Count("test", "t")) //python's count()
	p(s.Count("111111010101010", "1"))     //counts '1's

	p("Index:     ", s.Index("test is a test", "is"))                        //index - 5  python's find(), index()
	p("Index from last:     ", s.LastIndex("test is a test is hello", "is")) //index - 15

	p("Repeat:    ", s.Repeat("ab", 5)) //ababababab    similar to print('ab' * 5)

	//replace the string
	p("Replace:   ", s.Replace("fooooo", "o", "0", -1))    //replaces for all occurances
	p("Replace:   ", s.Replace("foooooo", "o", "0", 3))    //replaces only for "3" occurance
	p("ReplaceAll:   ", s.ReplaceAll("foooooo", "o", "#")) //replaceall defaultly

	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p(s.SplitN("a,b,c,d,e,f,g", ",", 4))             //["a" "b" "c" "d,e,f,g"]   //splits only first 4 elements
	p("Split:     ", s.SplitAfter("a-b-c-d-e", "-")) //[a- b- c- d- e] it splits with seperated string.
	// s.SplitAfterN()

	p("Join:      ", s.Join([]string{"a", "b"}, "-"))

	p("String_TRIM ", s.Trim("¡¡¡Hello, Gophers!!!", "!¡")) //python's strip()
	// s.TrimSpace()
	// s.TrimPrefix()
	// s.TrimSuffix()
	// s.TrimLeft()
	// s.TrimRight()

	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

	p("HasPrefix: ", s.HasPrefix("test", "te")) //startwith()
	p("HasSuffix: ", s.HasSuffix("test", "st")) //endwith()

	//converts string into string array using split

	fmt.Printf("%q\n", s.Split("sathooo", "")) //["s" "a" "t" "h" "o" "o" "o"]   //python list()

	st := "Have a nice day"
	str_array := (s.Split(st, " "))                                         //python split()
	p("Data type and Data -", reflect.TypeOf(str_array), "----", str_array) //data type - []string , Data - ["Have" "a" "nice" "day"]

	//converts string array into string using join
	join_str := (s.Join(str_array, " "))
	p(reflect.TypeOf(join_str), "----", join_str)

	p("strings are equal in upper/lower letters", s.EqualFold("sathish", "SATHISH"))

	//deep copy of a string
	oldString := "sathish"
	newString := s.Clone(oldString)
	newString = newString + "s"
	p(&oldString, oldString, &newString, newString) //different memory address and values

	//import unicode, Uses to check the ASCII values for is number, letter, upper,lower etc.
	unicode.IsNumber('9') //only characters/ascii/byte/rune allowed,
	unicode.IsLetter('E')
	unicode.IsSpace(' ')
	unicode.IsUpper('R')
	unicode.IsLower('r')

}
