package modules

import "fmt"

func LearnSwitch() {
	temp := 10

	switch temp {
	case 10:
		fmt.Println("Temp is 10 degrees")
	case 20:
		fmt.Println("Temp is 20 degrees")
	case 30:
		fmt.Println("Temp is 30 degrees")
	default:
		fmt.Println("Temp is unknown")
	}
}
