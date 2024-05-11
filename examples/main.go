package main

import (
	"fmt"

	"github.com/golangtime/measurer/length"
)

func main() {
	l, err := length.Parse("+10.2m")
	if err != nil {
		panic(err)
	}

	fmt.Println(l)

	l1 := 5 * length.Meter
	l2 := 5 * length.Kilometer

	fmt.Println(l1.Add(l2).Add(l).Meters())

	fmt.Println(length.Meter.Kilometers())
}
