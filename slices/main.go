package main

import "fmt"


func main()  {
	fmt.Println("Welcome to the Slices")

	var fruitsList= []string{"apple", "mango", "banana"};
	fmt.Println("fruit list item=", fruitsList)

	fruitsList=append(fruitsList, "tomato", "potato")  //we are adding more data in fruitslist slices.
	fmt.Println(fruitsList)

	fruitsList=append(fruitsList[1:]) ; //it means here we are adding all fruitlist item in fruitlist but from 1 index not 0;
	fmt.Println(fruitsList)

	fruitsList=append(fruitsList[1:3]) ; //it means here we are adding all fruitlist item in fruitlist but from 1 index not 0 to 3
	fmt.Println(fruitsList)
}