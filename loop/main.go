package main

import "fmt"


func main()  {
	fmt.Println("Welcome to the loop")

	days:=[]string{"Sunday", "Monday","Tuesday","Wednesday","Thousday","Friday","Saturday"};
	fmt.Println(days)


	//Normal for loops

	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	//range for loop

	for i:=range days{
		fmt.Println(days[i]) //for get one by one value.

		fmt.Printf("index is %v and value is %v", i,days[i])

		//This is break and continue statment are same like other programming language.
		if(i==2){
			break;
		}
	}


}