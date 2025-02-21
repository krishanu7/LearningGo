package modules

import (
	"fmt"
)	

func LearnForLoop() {
	// for i:=0; i<5; i++ {
	// 	fmt.Println("Value of i:", i);
	// }
	arr := []int{14,21,34,45,56};

	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value);
	}
}