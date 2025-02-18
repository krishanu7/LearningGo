package modules

import (
	"fmt"
	"sort"
)

var PublicStr string = "Public"

func Variables() {
	var x int = 200000
	y := x
	var dec float32 = 0.06
	var name string = "Krishanu"
	var isTrue bool = (name == "krishanu")
	var privateStr string = "Private"
	var persons [5]string
	persons[0] = "Krishanu"
	persons[1] = "Raj"
	arr := []int{2, 4}
	arr = append(arr, 1)
	arr = append(arr, 3)
	sort.Ints(arr)

	fmt.Printf("Integer x is %d, y is %d\n", x, y)
	fmt.Printf("Decimal is %f\n", dec)
	fmt.Printf("Name is %s\n", name)
	fmt.Printf("Is it true? %t\n", isTrue)
	fmt.Println("Private string is", privateStr)
	fmt.Println("Array is", arr)
	fmt.Println("Length of arr is", len(arr))
	fmt.Println("Array Element at pos 1 is", arr[1])
	fmt.Println("Persons are", persons)
}
