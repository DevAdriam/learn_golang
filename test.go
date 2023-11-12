package main

import "fmt" 

func main (){
	fmt.Println("hello world")

	// string and variables
	var name string = "mgmg"
	var age = 18
	var email string = "mgmg@gmail.com"
	var password string

	fmt.Println(name,email,password,age)

	// change value 
	name = "myamya"
	age =30

	fmt.Println(name,email,password,age)

	// initializing value using column instead of var
	location := "yangon"
	
	fmt.Println(location)

}

