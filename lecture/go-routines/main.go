package main

import "fmt"
import "time"

func main() {
	go sexyCount("test")
	go sexyCount("hoho")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
