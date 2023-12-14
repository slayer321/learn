//https://youtu.be/kaZOXRqFPCw?si=KOG8dmVqNEu9Tc_r

package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", 103120)
	userID := 100
	val, err := userData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result", val)
	fmt.Println("time took", time.Since(start))
}

type response struct {
	val int
	err error
}

func userData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	// val, err := fetchThirdPartStuffWhichCanBeShow(ctx)
	// if err != nil {
	// 	return 0, err
	// }
	respch := make(chan response)

	go func() {
		val, err := fetchThirdPartStuffWhichCanBeSlow()
		respch <- response{val, err}
		fmt.Println("Inside the go routine")
		fmt.Println(respch)
		valueOut := ctx.Value("userID")
		fmt.Println(valueOut)
	}()
	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("timeout")
	case resp := <-respch:
		return resp.val, resp.err
	}
	//return 12, nil
}

func fetchThirdPartStuffWhichCanBeSlow() (int, error) {

	time.Sleep(time.Millisecond * 50)
	return 666, nil

}
