package main

import "fmt"

func main()  {
	fmt.Println("welcome to the Pointers");

	// var ptr *int;
	// fmt.Println(ptr);

	myNumber :=3;

	var ptr = &myNumber; //here & means reffrence. means here we are getting reffrence of myNumber memory.

	fmt.Println("pointer of myNumber is =", ptr);
	fmt.Println("value of myNumber is =", *ptr);  //here we are getting value of this refference.

	*ptr= *ptr *2;
	fmt.Println("new value of my number is ", *ptr)

	
}