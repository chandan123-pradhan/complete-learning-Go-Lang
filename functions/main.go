package main

import "fmt"

func main()  {
	fmt.Println("Welcome to the functions");
	greeter()
	addTwoNumber(2,3)

	result:=addTwoNumberWithReturnType(3,4)
	fmt.Println("new result=",result)

	adderResult:=adder(2,3,4,5)  //here passing multiple values to add.
	fmt.Println(adderResult)


	name,age:=getMyNameAndAge();  //this function will return two params, name and age.

	fmt.Printf("my name is %v and age is %v",name,age)
}

func greeter(){
	fmt.Println("Hey I am Greeter Function")
}

func addTwoNumber(firstVal int,secondVal int){    //firstVal is value and int is datatype of this parameter.
	fmt.Println("result=", firstVal+secondVal);
}

func addTwoNumberWithReturnType(firstVal int,secondVal int)int{    //firstVal is value and int is datatype of this parameter.
	return firstVal+secondVal;
}


func adder(values ...int)int{    //here values are list of value and ...int means all values will be integer.
	total:=0;

	for _,val:= range values{
		total+=val;

	}
	return total;

}


func getMyNameAndAge()(string,int){
return "chandan",23
}