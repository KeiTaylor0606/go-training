package advancedgrammer

import "fmt"

func total(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t
}

func channel() {
	c := make(chan int)
	go total(100, c)
	fmt.Println("total:", <-c)
}

func channel_2() {
	c := make(chan int)
	go total(1000, c)
	go total(100, c)
	go total(10, c)
	fmt.Println("total:", <-c, <-c, <-c)
}
