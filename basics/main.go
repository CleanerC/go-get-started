package main

import "fmt"

var b string

func add(a int, b int) int { 		//add(a, b int)
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("Hello, World!!")

	a := 1
	fmt.Println(a)


	b = "This is B"
	fmt.Printf("B is typed: %T  string: %s", b, b)

	c := add(1, 3)
	fmt.Println(c)

	b := 9
	a, b = swap(a, b)
	fmt.Printf("a is now: %d, b is now %d", a, b)

	list1 := [4]int{1, 2, 3, 4}
	for ii := 0; ii < 4; ii++ {
		fmt.Println(ii, "th item in list1 is: ", list1[ii])
	}

	slice1 := make([]int, 0)
	slice1 = append(slice1, 5, 4, 3, 2, 1)
	for ii := 0; ii < 5; ii++ {
		fmt.Println(ii, "th item in slice1 is: ", slice1[ii])		
	}

	map1 := map[string]int {  		//map[key]value
		"a": 1,
		"b": 2,
	}
	fmt.Println(map1)

	map2 := make(map[string]int) 
	map2["a"] = 1
	map2["c"] = 2
	fmt.Println(map2)


}

