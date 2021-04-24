package main

import (
	"fmt"

	"github.com/saharak-manoo/etracking-sdk-go/etracking"
)

func main() {
	etrackingApi, err := etracking.New("YOUR_API_KEY", "YOUR_KEY_SECRET")
	if err != nil {
		fmt.Println("Errors -> ", err.Error())
	}

	// Find courier = jt express, tracking number = 860050536402
	fmt.Println("Find courier = jt express, tracking number = 860050536402")
	fmt.Println(etrackingApi.Find("jt-express", "860050536402"))

	// Find Jt express tracking number = 860050536402
	fmt.Println("Find Jt express tracking number = 860050536402")
	fmt.Println(etrackingApi.JtExpress("860050536402"))
}
