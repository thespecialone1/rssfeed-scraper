package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	fmt.Println("Hey yo")
	portString:= os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviorment")
	} else{
		fmt.Println("PORT IS:", portString)
	}
}
