package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Param1 string `json:"username"`
}

func main() {
	fmt.Print("Hello World \n")
	users := []*User{
		{"vasily"},
		{"ivan"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Hello Test \n")
		resp, _ := json.Marshal(users)
		w.Write(resp)
		fmt.Print(resp)
	})
	http.ListenAndServe(":8081", nil)
}
