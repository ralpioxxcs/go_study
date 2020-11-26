package main

import "fmt"
import "time"

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	//go sexyCount("test")
	//go sexyCount("hoho")
	//time.Sleep(time.Second * 10)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
