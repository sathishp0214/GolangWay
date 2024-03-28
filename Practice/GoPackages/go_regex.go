package main

import (
	"fmt"
	"regexp"
)

/*
Basic Regex Expressions:

^ - start of the string
$ - End of the string(once gave, doesn’t take any parameters after this $)
\s - only spaces
\S - except spaces, match alphabets, symbols and numbers
\d - only numbers
\D - except numbers, match  alphabets, numbers,spaces and symbols
\w - only alphabets(both uppercase and lowercase) and numbers
\W - only symbols(@#%&./ etc) and spaces


\w{2,3}    - allows only alphabets (at least two letters to maximum three letters)
\w{3,}  - atleast three letters with “N” number of output limit


Multi-letter matching : (+) symbol should add at the suffix

\w+  - alphabets-numbers combinations (ex:ffg33,44gg,dd44aa,gfrhhh,14646)
\W+  -  Ex: $#%@ , /.%
\d+  - Ex;356,32213

Rub+ty  - Rubbbbbbbbbty  # (+) symbol will match multiple ‘b’ in the middle as well.


|  - OR operator - Ex:^192.168.([0-9]{2}|[0-9]{3}).([0-9]{2}|[0-9]{3})$ #() bracket should use between | OR operator.

And operator - Did not find

\. - for dots

Use of (  )   #creates groups in the regex
Use of +   #[a-zA-Z0-9]+   #multi letter matching

(.*) vs (.*?):

(.*)  - Takes everything upto end of string. Ex: abcdefg  -  ab(.*)

(.*?) - Takes Everything in between/beginning of string.
Ex: abcdefg  -  ab(.*?)fg  #it takes everything between.
Ex: abcdefg  -  (.*?)defg  #it takes everything in beginning.


 * vs +:

s* -  matches zero value or multiple ‘s’ value
s+ - matches at least one value of ‘s’ or multiple ‘s’ value




[a-z], [A-Z], [0-9]  - Matches lower case, UpperCase and numbers respectively.
([a-z]{4}) , ([0-9]{3}) - Matches continuous 4-digit lowercase letters.
[wsert]+ - Matches any of from these letters
[^a-z] ,  [^A-Z], [^0-9] - Matches any other than  lower case, UpperCase and numbers respectively.
[Pp]ython - Match "Python" or "python"
A(nt|pple) - matches Apple or Ant

rub(y|le) - Match "ruby" or "ruble"
He..o - Matches Hello
[a-zA-Z0-9]  -  Matches alphabets and numbers
A*B*C*  -  AAAAAABBBC    #matches same letter None or multiple times
Https? - matches Http or Https
\[\{\(\)\}\]  - [{()}]       otherwise  [{()}]+  -  [{()}]
\t - matches TAB button only   ex:T\t\w{2}   -  T	ab
A?  - matches empty or A

(?i)teST  - Matches TEST,tEsT,    #Any Case match - ignores case insensitive
*/

func main() {

	// st := "sathish thi"
	// regex := "this"

	//matches date format
	// st := "If I am 20 years 10 months and 14 days old as of August 17,2016 then my DOB would be 1995-10-03"
	// regex := `\d{4}-\d{2}-\d{2}`    //use ` ` for raws string - python similar to r" "

	//validates email address
	st := "sathoo@sat.com"
	regex := `\w+@\w+.com`

	regex_compile := regexp.MustCompile(regex) //regexp.MustCompile() - In golang always use this for more convenience   similar to python re.compile()

	fmt.Println(regex_compile.MatchString(st)) //true (or) false   //similar to python re.match()

	fmt.Println(regex_compile.FindString(st)) //prints finded string for regex  ex: "1995-10-03", If Doesn't find prints "blank"

	fmt.Println(regex_compile.FindStringIndex(st)) //[2 6]   //regex matched index range in input string

	fmt.Println(regex_compile.ReplaceAllString(st, "#####")) //replaces similar to python re.sub()

	//FindString  vs FindAllString

	st1 := "sathish this"
	regex1 := "this"

	regex_compile = regexp.MustCompile(regex1)
	fmt.Println(regex_compile.FindAllString(st1, -1))      //[this this]   FindString - just match only first "this" and stop
	fmt.Println(regex_compile.FindAllStringIndex(st1, -1)) //[[2 6] [8 12]] - gets all match indexes

	//FindString    - Finds only one match  //similar to python re.find()
	//FindAllString (Mostly use this) - Finds all matches in string (or) multiple string  //similar to python re.findall()

	regexValues()

}

func regexValues() {
	v := `world 45 hello 5677 world 345
	fdg djd 467 hello dgf 546 hello`
	r := regexp.MustCompile(`(\d+) hello`)
	// match := r.MatchString(v)   //return true or false if match found or not
	// match := r.FindString(v) //returns first matched value only
	// match := r.FindAllString(v, -1) //returns all matches continuously in slice, -1 means all match occurance
	// match := r.ReplaceAllString(v, `#######`) //Replace all matches with #######
	// match := r.FindStringSubmatch(v) //useful to fetch the first matching regex group values and returns in slice
	match := r.FindAllStringSubmatch(v, -1) //all version of above submatch

	fmt.Println(match)

}
