package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Structs")

	//Note:- There is no Inheritance in go lang and no super and parent.

	user := User{"chandan", "chandan@go.dev", true, 32}
	fmt.Println(user)
	fmt.Printf("chandan's details= %+v",user)  // output:-chandan's details= {Name:chandan Email:chandan@go.dev IsVerified:true Age:32}

	//fmt.Println(user.Name) //to get any specific item;

}

//Now creating a struct.

type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
}
