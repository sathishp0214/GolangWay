package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			fmt.Printf("%+v===============%+v=============%+v", current, current.next, data)
			fmt.Println()
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Display() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	// list := LinkedList{}

	// list.Insert(10)
	// list.Insert(20)
	// list.Insert(30)
	// list.Insert(40)

	// list.Display()

	//https://blog.boot.dev/golang/merge-sort-golang/
	//merge sort uses divide and conquer algorithm,
	//merge sort logic:
	//divide: Divides the input slice until upto small length of 1 using recursion,
	//conquer: Then sorts the all divided slices sepertely and keep on appends into final slice. Finally gets sorted slice.
	fmt.Println("sorted slice---", mergeSort([]int{10, 6, 2, 1, 12, 5, 8, 3, 11, 4, 7, 9}))
}

func mergeSort(InputSlice []int) []int {
	if len(InputSlice) < 2 { //if len of slice 0 or 1, then returns it
		return InputSlice
	}

	first := mergeSort(InputSlice[:len(InputSlice)/2])
	second := mergeSort(InputSlice[len(InputSlice)/2:])

	fmt.Println(InputSlice, "slices---------", first, "=============", second)

	return merge(first, second)
}

func merge(first []int, second []int) []int {
	final := []int{}
	i := 0
	j := 0

	//from the two slices, sorting the values upto the equal length
	for i < len(first) && j < len(second) {
		if first[i] < second[j] { //Here doing ascending order, If descending order have to change into ">"
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}

	fmt.Println(i, j, "=============", first, "===============", second, "---------before-----", final)

	//whatever first and slice values are yet to append into final slice, Here appending it
	for ; i < len(first); i++ { //variation of for loop without "i" initialization value
		final = append(final, first[i])
	}

	for ; j < len(second); j++ {
		final = append(final, second[j])
	}

	fmt.Println(i, j, "=============", first, "===============", second, "---------after-----", final)

	return final

}
