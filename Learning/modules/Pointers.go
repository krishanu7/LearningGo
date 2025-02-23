package modules

import ("fmt")

type person struct {
	Name string
	Age int
}

func updatePerson(p *person) {
	p.Name = "Krishanu Saha"
	p.Age = 22
}

func magicPointer(ptr *int){
	*ptr = 3;
}

func Pointers() {
	// var val int;
	// val = 2;
	// var ptr *int ;
	// ptr = &val;
	val := 2;
	ptr := &val;
	fmt.Println("Value of val is ", val);
	fmt.Println("Value of ptr is ", ptr);
	fmt.Println("Value of *ptr is ", *ptr);
	magicPointer(&val);
	fmt.Println("Value of val is ", val);
	
	p := person{
		Name: "Krishanu",
		Age: 24,
	}
	fmt.Println("Person is ", p);
	updatePerson(&p);
	fmt.Println("Person is ", p);
	arr := [5]int{1,2,3};
	nums := &arr;
	fmt.Println("Array is ", arr);
	fmt.Println("Array is ", nums[0]); //Auto-Dereferencing  works in array but not in slice
	fmt.Println("Array is ", (*nums)[0]);
}