package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	welcome :="welcome to the user input";
	fmt.Println(welcome);

	relader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter The Rating: ");

	// comma ok || comma err syntax
	// input, err:=relader.ReadString('\n');
	input, _:=relader.ReadString('\n');
	fmt.Println("Thanks for the rating", input);
	//fmt.Println("error is occured", err);//if there is any error occured.

	//conversion.
	numRating,error :=strconv.ParseFloat(strings.TrimSpace(input),64);

	if(error!=nil){
		fmt.Println(error);
	} else {
		fmt.Println("Added 1 to your rating :", numRating+1);
	}
	
	

}