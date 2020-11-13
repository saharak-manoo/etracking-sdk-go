package main

import (
	"fmt"

	"github.com/saharak/etracking-sdk-go/etracking"
)

func main() {
	// etrackingApi, err := etracking.New("<etrackings api key>", "<etrackings key secret>")
	etrackingApi, err := etracking.New("899de77a545537ac1561f277ec57fdb49610bb5a", "b5086e7c36d7936d0caa954e2a836917829bc56b322a91981e019a98b84dd60b9d9a493e401366cf91d1a0e7212daf33f99b1da20daaf150aefef10644b2c71a37baab196c691ebfd3c680")
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