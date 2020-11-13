package main

import (
	"fmt"

	"github.com/saharak/etracking-sdk-go/etracking"
)

func main() {
	etrackingApi, err := etracking.New("<etrackings api key>", "<etrackings key secret>")
	if err != nil {
		fmt.Println("Errors " + err.Error())
	}

	// Find courier = jt express, tracking number = 860050536402
	fmt.Println("Find courier = jt express, tracking number = 860050536402")
	fmt.Println(etrackingApi.Find("jt_express", "860050536402"))

	// Find Jt express tracking number = 860050536402
	fmt.Println("Find Jt express tracking number = 860050536402")
	fmt.Println(etrackingApi.JtExpress("860050536402"))
}