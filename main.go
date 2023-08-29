// following AnthonyGG YT vid, youtube.com/watch?v=kaZOXRqFPCw

package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err error
}

func main() {
	start := time.Now()
	userID := 10

	ctx := context.Background()
	

	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: ", val)
	fmt.Println("Time taken: ", time.Since(start))
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 200);
	defer cancel()
	respch := make(chan Response)

	go func() {
		val, err := fetchSlowService()
		respch <- Response{
			value: val,
			err: err,
		}
	}()

	for {
		select {
		case <- ctx.Done():
			return 0, fmt.Errorf("fetching from 3rd party took too long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchSlowService() (int, error) {
	time.Sleep(time.Millisecond * 500)
	return 666, nil
}