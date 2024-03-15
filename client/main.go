package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Response struct {
	Status Status `json:"status"`
}

func showStatus() {
	res, err := http.Get("http://localhost:8080/status")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyString := string(body)

	var response Response
	err = json.Unmarshal([]byte(bodyString), &response)

	status := response.Status

	if err != nil {
		fmt.Println(err)
	}

	waterCondition := ""
	if status.Water <= 5 {
		waterCondition = "aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		waterCondition = "siaga"
	} else {
		waterCondition = "bahaya"
	}

	windCondition := ""
	if status.Wind <= 6 {
		windCondition = "aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		windCondition = "siaga"
	} else {
		windCondition = "bahaya"
	}

	fmt.Println("=========================")
	fmt.Println(time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Println("-------------------------")
	fmt.Printf("Ketinggian air : %vm\n", status.Water)
	fmt.Printf("Status air : %s\n", waterCondition)
	fmt.Println("-------------------------")
	fmt.Printf("Kecepatan angin : %vm/s\n", status.Wind)
	fmt.Printf("Status angin : %s\n", windCondition)
	fmt.Println("=========================")
}

func main() {
	fmt.Println("response will be displayed and updated every 15 second")

	for range time.Tick(time.Second * 15) {
		showStatus()
	}
}
