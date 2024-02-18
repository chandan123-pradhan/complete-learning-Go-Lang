package main

import "fmt"

func main()  {
	fmt.Println("Welcome to the methods")

	user:=User{"chandan",23}

	user.GetUserDetails()
	user.NewName()
	user.setNewName("dipu kumar")
}



type User struct{
	Name string;
	Age int;

}

func(u User) GetUserDetails(){
	fmt.Printf("user's name is %v and age is %v", u.Name,u.Age);
}

func(u User) NewName(){
	u.Name="kaushiki kumari";
	fmt.Println("new user name is ", u.Name)
}


func(u User) setNewName(newName string){
	u.Name=newName
	fmt.Println("new user name is ", u.Name)
}