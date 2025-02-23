package modules

import (
	"fmt"
	"strconv"
)

func Conversion() {
	var x int = 10
	var y float64 = float64(x)
	var uintVar uint = uint(x)
	fmt.Println("Int:", x)
	fmt.Println("Int to Float:", y)
	fmt.Println("Int to Uint:", uintVar)
	var str1 string = strconv.Itoa(x)
	fmt.Println("Int to String:", str1)
	var str2 string = "30"
	var intVar, _ = strconv.Atoi(str2)
	fmt.Println("String to Int:", intVar)
	floatNum := 12.34
	str3 := strconv.FormatFloat(floatNum, 'f', 2, 64)
	fmt.Println("Float to String:", str3)
	fmt.Printf("Type of str3: %T\n", str3)
	floatNum, _ = strconv.ParseFloat(str3, 64)
	fmt.Println("String to Float:", floatNum)
	fmt.Printf("Type of floatNum: %T\n", floatNum)

	str4 := "Hello"
	byteSlice := []byte(str4) // Convert string to byte slice
	fmt.Println("Byte Slice:", byteSlice)

	/*
		Conversion	    Function
		Int → Float	    float64(intVar)
		Float → Int	    int(floatVar)
		Int → String	strconv.Itoa(intVar)
		String → Int	strconv.Atoi(strVar)
		Float → String	strconv.FormatFloat(floatVar, 'f', 2, 64)
		String → Float	strconv.ParseFloat(strVar, 64)
		String → ByteSlice	[]byte(strVar)
		ByteSlice → String	string(byteSlice)

	*/
}
