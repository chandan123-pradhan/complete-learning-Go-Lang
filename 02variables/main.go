package main

import "fmt"

const LoginToken="sdfsdfsersg";  //declared Public constant variable.

func main()  {
	var username string ="Chandan Pradhan";
	fmt.Println(username)
	fmt.Printf("Variable Type: %T \n", username) //It will print The data type of the variables.


	var isLoggedIn bool =true;
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable Type: %T \n", isLoggedIn)  //It will print The data type of the variables.

	var age int =23;
	fmt.Println(age);
	fmt.Printf("variable type = %T \n",age); //it will print the data type of the variable.


	//default value  and some alias
	var defaultVal int;
	fmt.Println(defaultVal);  //without initialise any value it will store 0 in int type.
	fmt.Printf("variable type = %T \n",defaultVal); //it will print the data type of the variable.

	//implicit Type
	var name="chandan pradhan";
	fmt.Println(name);

	// no var variables.
	numberOfGf :=4;
	fmt.Println(numberOfGf)


	fmt.Printf("Login Token= %T \n",LoginToken);




}