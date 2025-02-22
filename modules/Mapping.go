package modules

import ("fmt")

func MapinGo () {
	// way1 to declare map
	// myMap := make(map[string]int);
	// myMap["one"] = 1;
	// myMap["two"] = 2;
	// myMap["three"] = 3;
	// fmt.Println(myMap["one"]);
	// way2 to declare map
	myMap := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	}

	value, exist := myMap["four"];
	if exist {
		fmt.Println(value);
	}
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value);
	}
	delete(myMap, "two");
	fmt.Println("----------------- After deleting two -----------------");
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value);
	}
	//Nested map
	nestedMap := map[string]map[string]string {
		"krishanu":  {
			"email": "krishau.s@gmail.com",
			"phone": "1234567890",
		},
		"bipasha": {
			"email": "bipasha.b@gmail.com",
			"phone": "0987654321",
		},
	}
	fmt.Println("----------------- Nested Map -----------------");
	fmt.Println("Nested Map:", nestedMap["bipasha"]["email"]);
}