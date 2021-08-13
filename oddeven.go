package main

import (
	"fmt"
	"sync"
)

//WaitGroup
var wg = sync.WaitGroup{}

//Odd will print all Odd numbers inside 100
//taking parameters as channels to communicate
// with other goroutines
func Odd(c chan bool) {
	var value []int
	for i := 0; i <= 100; i++ {
		if i%2 == 1 {
			<-c
			fmt.Println(i)
			value = append(value, i)
			c <- true
		}
	}
	wg.Done()
}

//Even will print all even numbers inside 100
//taking parameters as channels to communicate
// with other goroutines
func Even(c chan bool) {
	var value []int
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			<-c
			fmt.Println(i)
			value = append(value, i)
			c <- true
		}
	}
	wg.Done()
}

func main() {
	c := make(chan bool, 1)
	defer close(c)
	wg.Add(2)

	go Odd(c)
	go Even(c)
	c <- true
	wg.Wait()

}
