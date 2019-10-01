package main

import (
	"fmt"
)

func main() {
	te := Remo("https://api.nature.global/1/devices")
	fmt.Printf("%f\n", te)
	PushData(te)
}
