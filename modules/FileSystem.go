package modules

import (
	"fmt"
	"os"
)

func FileSystem() {
	file, err := os.Create("example.txt")

	if err!=nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("File created successfully")
}
