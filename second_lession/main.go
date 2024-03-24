package main

import (
	"fmt"
	"github.com/evgeniy-p/go_learn/second_lession/logic"
	"os"
)

func main() {
	actions := []string{"осмотреться", "идти коридор", "идти комната", "осмотреться", "надеть рюкзак", "взять ключи", "взять конспекты",
		"идти коридор", "применить ключи дверь"}
	data, rooms, err := logic.InitGame("New Player")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(data.FirstAction, data.RoomNow)
	//fmt.Println(rooms)
	for _, action := range actions {
		returnedText, err := logic.HandleCommand(action, data, rooms)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(returnedText)
	}

}
