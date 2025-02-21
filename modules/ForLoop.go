package modules

import (
	"fmt"
)

func LearnForLoop() {
	// for i:=0; i<5; i++ {
	// 	fmt.Println("Value of i:", i);
	// }
	arr := []int{14, 21, 34, 45, 56}

	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}
	var i = 0
	for {
		fmt.Println("Value of i:", i)
		i++
		if i == 5 {
			break
		}
	}
}
