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

	//Number(ints)
	var numberOne int = 1
	var numberTwo int = 2
	numberThree := 3
	
	fmt.Println(numberOne,numberThree,numberTwo)


	// bits & memory
	var numFive int8 = 127  
	var numSix int16 = 8989

	// unsigned int means(only accept positive int)
	var numSeven uint8 = 3232
	fmt.Println(numFive,numSix,numSeven)

}

