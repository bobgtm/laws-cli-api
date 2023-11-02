package main

import (
	"fmt"
	"strconv"
)

// Per the LegiSCan API User Manual
// Most use cases for the api can be satsified with daily updates
// although recommended frequency varies for different operations
func callBackBill(cfg *config, args ...string) error {
	bills, err := cfg.lawsClient.ListBills()
	if err != nil {
		return fmt.Errorf("error fetching Bills object: %v", err)
	}
	for i := 0; i < 5; i++ {
		number := bills.Masterlist[strconv.Itoa(i)].Number
		title := bills.Masterlist[strconv.Itoa(i)].Title
		// statDate := bills.Masterlist[strconv.Itoa(i)].StatusDate
		// status := bills.Masterlist[strconv.Itoa(i)].Status
		lastAct := bills.Masterlist[strconv.Itoa(i)].LastAction
		lastActDate := bills.Masterlist[strconv.Itoa(i)].LastActionDate
		desc := bills.Masterlist[strconv.Itoa(i)].Description

		fmt.Printf("%v - %v: %v\n", lastActDate, number, title)
		fmt.Printf("Description > %v\n", desc)
		fmt.Printf("Last Action > %v\n", lastAct)

		fmt.Println("-------------")
	}
	return nil
}
