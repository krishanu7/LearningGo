package modules

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	// var x int;
	// var y int;
	// var z int;

	// fmt.Printf(("Enter the first number: "));
	// fmt.Scan(&x);
	// fmt.Printf(("\nEnter the second number: "));
	// fmt.Scan(&y);
	// z = x + y;
	// fmt.Printf("\nSum of %d and %d is %d\n", x, y, z);

	fmt.Println("Hey, what's your name?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Printf("Hello, %s", name)
}
