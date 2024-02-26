package main

import "fmt"


func main()  {
	
// when you'll use defer keyword before any line, it means this line will be execute by end of the function.
defer fmt.Println("world")
defer fmt.Println("Third")
defer fmt.Println("Two")
defer fmt.Println("One")
fmt.Println("Hello")
myDifer();
//output:- Hello World.

//In the case of mulitple defer it will run in LIFO(last in first out) order.
}

func myDifer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}

	//output is 4,3,2,1,0 .
}