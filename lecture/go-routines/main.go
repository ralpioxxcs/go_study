package main

import "fmt"
import "time"

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "thor", "peter", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for", i)
		fmt.Println(<-c) // blocking operation,, waiting channel
	}

	//resultOne := <-c
	//resultTwo := <-c
	//resultThree := <-c

	//fmt.Println("Waiting for message")
	//fmt.Println("Received this message:", resultOne)
	//fmt.Println("Received this message:", resultTwo)
	//fmt.Println("Received this message:", resultThree) // runtime error!
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
