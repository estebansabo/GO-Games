package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("BUENOS DIAS")
	} else if hora < 17 {
		fmt.Println("BUENAS TARDES")
	} else {
		fmt.Println("BUENAS NOCHES")
	}
}
