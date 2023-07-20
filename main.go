package main

import (
	"fmt"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("BUENOS DIAS")
	} else if t.Hour() < 17 {
		fmt.Println("BUENAS TARDES")
	} else {
		fmt.Println("BUENAS NOCHES")
	}
}
