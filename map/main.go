package main

import "fmt"


func main()  {
	fmt.Println("Welcome to the map");

	languageList :=make(map[string]string);
	
	languageList["js"]="Javascript";
	languageList["ds"]="Dart";

	fmt.Println("langugage list= ", languageList)

	fmt.Println("java script=", languageList["js"])

	// delete(languageList,"js")  //for delete specific key.
	// fmt.Println(languageList)


	//Loops

	for key,value:=range languageList{
		fmt.Println("key is = %v, value is= %v\n",key,value)
	}

}

