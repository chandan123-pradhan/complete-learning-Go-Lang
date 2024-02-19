package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


func main()  {
	fmt.Println("welcome to the files handing in go")

	content:="Hello This content is generated from go program";

	file, error := os.Create("./myAutoCreatedFile.txt")  //here we are creating file with file name and data types.

	checkNilError(error)  //if there is any error just program will be go into this function and program will be end.
	length, err := io.WriteString(file,content);  //here we are writing content into created file.
	if(err!=nil){
		panic(error)
	}
	fmt.Println("length is =", length)
	defer file.Close()  //it means here file are closed.
	readFile("./myAutoCreatedFile.txt")
}


func readFile(fileName string)  {
	dataByte, err :=ioutil.ReadFile(fileName);  //reading that file.

	checkNilError(err)  //if there is any error just program will be go into this function and program will be end.

	stringData :=string(dataByte)  //reader method return into dataByte formate so here we are converting into string format,

	fmt.Println("content of the file is \n", stringData);

}


func checkNilError(err error)  {
	if err!=nil{
		panic(err);
	}
}