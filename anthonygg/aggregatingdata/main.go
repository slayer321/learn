package main

import (
	"fmt"
	"time"
)

// func main() {
// 	start := time.Now()
// 	userName := fetchUser()
// 	likes := fetchLikes(userName)
// 	match := fetchMatchUser(userName)

// 	fmt.Println("likes:", likes)
// 	fmt.Println("match:", match)
// 	fmt.Println("took:", time.Since(start))
// }

// func fetchUser() string {
// 	time.Sleep(time.Millisecond * 100)

// 	return "BOB"
// }

// func fetchLikes(userName string) int {
// 	time.Sleep(time.Millisecond * 150)

// 	return 11
// }

// func fetchMatchUser(userName string) string {
// 	time.Sleep(time.Millisecond * 100)

// 	return "ANNA"
// }

// func main() {
// 	start := time.Now()
// 	cLikes := make(chan int)
// 	cUser := make(chan string)
// 	cMatch := make(chan string)
// 	go fetchUser(cUser)
// 	go fetchLikes(cUser, cLikes)
// 	go fetchMatchUser(cUser, cMatch)

// 	fmt.Println("likes:", <-cLikes)
// 	fmt.Println("match:", <-cMatch)
// 	fmt.Println("took:", time.Since(start))
// }

// func fetchUser(cUser chan string) chan string {
// 	//time.Sleep(time.Millisecond * 100)
// 	cUser <- "BOB"
// 	return cUser
// }

// func fetchLikes(userName chan string, likes chan int) chan int {
// 	//time.Sleep(time.Millisecond * 150)
// 	likes <- 11
// 	return likes
// }

// func fetchMatchUser(userName chan string, cMatch chan string) string {
// 	//time.Sleep(time.Millisecond * 100)
// 	cMatch <- "ANNA"
// 	return <-cMatch
// }

func main() {
	start := time.Now()

	respCh := make(chan any, 2)
	//wg := sync.WaitGroup{}

	//wg.Add(2)

	userName := fetchUser()
	resp := make([]any, 2)
	go fetchLikes(userName, respCh)
	go fetchMatchUser(userName, respCh)

	//close(respCh)
	//for resp := range respCh {
	// fmt.Printf("resp: %v\n", <-resp)
	// for i := range respCh {
	// 	respCh[i] <- respCh
	// }
	//}
	for i := range resp {
		gg := <-respCh
		//fmt.Println(gg)
		resp[i] = gg
	}
	for _, val := range resp {
		fmt.Println("out", val)
	}
	fmt.Println("took:", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

func fetchLikes(userName string, respCh chan any) {
	time.Sleep(time.Millisecond * 150)

	respCh <- 11
}

func fetchMatchUser(userName string, respCh chan any) {
	time.Sleep(time.Millisecond * 100)
	respCh <- "ANNA"

}
