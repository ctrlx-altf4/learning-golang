package main

import (
	"fmt"
	"learningGo/randomGreeting/greeting"
)

func main(){
	message,err := greeting.Hello("Prajwal");

	if err != nil{
		panic(err)
	}


	fmt.Println(message);
}