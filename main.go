package main

import (
	"fmt"
	"strings"
)

type tempAndTime struct {
	Temp float64
	Time string
}

func main() {

	data, err := FetchData()
	if err != nil {
		fmt.Printf("error: %v", err)
		panic(err)
	}

	tempTime := []tempAndTime{}

	fmt.Printf("\n    Today it's\n\n")

	for _, strt := range data.List {
		if strings.Contains(strt.DtTxt, "00:00:00") {
			break
		}
		tempTime = append(tempTime, tempAndTime{Temp: strt.Main.Temp, Time: strt.DtTxt})
	}

	for _, val := range tempTime {
		fmt.Printf(" -> %v'C at %v\n", val.Temp, val.Time[11:16])
	}
}
