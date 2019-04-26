package main

import (
	"fmt"
	"time"
)

func diff() int64 {
	t0 := time.Now()
	for {
		t := time.Now()
		nano := t.Sub(t0).Nanoseconds()
		if nano != 0 {
			return nano
		}
	}
}

func warmup() {
	t0 := time.Now()
	for i := 0; i < 1000000; i++ {
		t := time.Now()
		nano := t.Sub(t0).Nanoseconds()
		if nano != 0 {
		}
	}
}

func main() {
	// warm up
	warmup()

	// measure and print the measurement overhead (measurement resolution)
	t0 := time.Now()
	for i := 0; i < 1000000; i++ {
		t := time.Now()
		nano := t.Sub(t0).Nanoseconds()
		if nano != 0 {
		}
	}
	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
	}
	t2 := time.Now()
	fmt.Print("Measurement Resolution: ")
	fmt.Println((t1.Sub(t0).Nanoseconds() - t2.Sub(t1).Nanoseconds()) / 1000000)	

	// measure and print the time resolution
	fmt.Print("Time Resolution: ")
	for i := 0; i < 10; i++ {
		fmt.Print(diff(), " ")
	}
	fmt.Print("\n")
}
