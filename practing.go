package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var tp = fmt.Println

func main() {
	a := 12
	var a1 int
	var a2 = 20
	tp(a, a1, a2)

	var y *int
	y = &a1
	tp(y, *y)

	g := []int{2, 3, 4}
	g = append(g, 10, 23, 45)

	g1 := [3]int{2, 4, 65}

	var u *[3]int
	u = &g1
	tp(u, *u)

	j := map[int]int{11: 11, 21: 22, 31: 33}

	delete(j, 100)
	tp(j, &j)

	for i := 0; i < len(g); i++ {
		tp(g[i] % 2)
	}

	for i := 0; i < len(g); i++ {
		for j := 0; j < i; j++ {
			if g[i] < g[j] {
				g[i], g[j] = g[j], g[i]
			}
		}
	}

	tp("after sorted", g)

	for i, k := range j {
		tp(i, k, j[i], j[k])
	}

	a = 0
	if a == 0 {

	} else if a < 0 {

	} else if a > 0 {
		tp("ghgg")
	} else {
		tp("yyyyyy")
	}

	if (a == 0 || a < 0) && a < 0 {
		tp("777777777")
	}

	a = 2
	switch a {
	case 0:
		tp("zero")
		fallthrough
	case 1:
		tp("one")
	default:
		tp("wrong")
	}

	var p bool
	p = true
	a = 10

	switch {
	case p == true && a == 10:
		tp("case 1")
		fallthrough
	case a > 5:
		tp("case 2")
		fallthrough
	default:
		tp("ffff")
	}

	var k map[int]interface{}
	k = map[int]interface{}{}
	k[1] = 304
	k[2] = "fff"
	k[3] = true
	k[4] = []int{3, 4, 5, 6}
	k[5] = 20.5
	k[6] = struct{}{}
	k[7] = struct{}{}

	// for i, j := range k {
	// 	switch j.(type) {
	// 	case []int:
	// 		tp(i, j)
	// 	case bool:
	// 		tp(i, j)
	// 	case float64:
	// 		tp(i, j)
	// 		// default:
	// 	}
	// }

	for i, j := range k {
		switch reflect.TypeOf(j).Kind() {
		case reflect.Slice:
			tp(i, j)
		case reflect.Bool:
			tp(i, j)
		case reflect.Struct:
			tp(i, j)
			// default:
		}
	}

	func() {
		tp(reflect.TypeOf(k), reflect.TypeOf(k).Kind())
	}()

	p1 := " 012345AZaz"
	o := []byte(p1)
	o1 := []rune(p1)
	tp(o, o1)

	yp := string(o)
	tp(yp)

	rt := []int{8, 43, 7, 769, 9, 0}
	rt1 := []int{345, 457}
	rt1 = append(rt1, rt...)
	tp(rt, rt1)

	gp := 10
	t1 := strconv.Itoa(gp)
	tp(t1, reflect.TypeOf(t1))

	hj := "102.567"
	ll, _ := strconv.ParseFloat(hj, 64)
	fmt.Println(reflect.TypeOf(ll), ll, int(ll))

}
