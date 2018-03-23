package main

import (
	"MasteringGoTutorial/ClassFactoryTutorial/appliances"
	"fmt"
)

func main(){
	fmt.Println("Enter prefered applianxe type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	var myType int
	fmt.Scan(&myType)

	myAppliance, err := appliances.CreateAppliance(myType)

	if err == nil{
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}

}