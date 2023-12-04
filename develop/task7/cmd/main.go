package main

import (
	"fmt"
	"task7/pkg"
	"time"
)

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-pkg.Or(sig(2*time.Second), sig(5*time.Second), sig(1*time.Second), sig(1*time.Second), sig(1*time.Second))

	fmt.Printf("fone after %v\n", time.Since(start))
}
