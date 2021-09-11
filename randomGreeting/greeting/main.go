package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string)(string,error){
	if name == ""{
		return name,errors.New("empty name")
	} else {
		message := fmt.Sprintf(randomFormat() ,name)
		return message,nil;
	}

}

func init(){
	rand.Seed(time.Now().UnixMilli())
}

func randomFormat()string{
	greetingFormats := []string{
		"Hola, %v",
		"Hello, %v",
		"Namaste, %v",
		"Hi, %v",
	}
	return greetingFormats[rand.Intn(len(greetingFormats))]
}