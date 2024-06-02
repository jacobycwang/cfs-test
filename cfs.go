package main

import (
	"crypto/sha512"
	"flag"
	"log"
	"syscall"
	"time"
)

func main() {
	burnInterval := flag.Duration("burn", time.Millisecond*5, "burn time")
	sleepInterval := flag.Duration("sleep", time.Second, "sleep between iterations")
	iterations := flag.Int("iterations", 10, "number of iterations")
	flag.Parse()

	time.Sleep(time.Second)

	b := time.Now()

	for i := 0; i < *iterations; i++ {
		s := time.Now()
		burn(*burnInterval)
		e := time.Now().Sub(s)

		log.Printf("[%d] burn took %dms, real time so far: %dms, cpu time so far: %dms", i, ms(e), ms(time.Since(b)), ms(usage()))

		time.Sleep(*sleepInterval)
	}
}

func ms(duration time.Duration) int {
	return int(duration.Nanoseconds() / 1000 / 1000)
}

func burn(duration time.Duration) {
	// s := time.Now()
	s := usage()

	for {
		sum := sha512.New()
		sum.Write([]byte("banana"))
		sum.Sum([]byte{})

		// if time.Since(s) > duration {
		// 	break
		// }
		if usage()-s > duration {
			break
		}
	}
}

func usage() time.Duration {
	r := syscall.Rusage{}
	syscall.Getrusage(syscall.RUSAGE_SELF, &r)
	return time.Duration(r.Stime.Nano() + r.Utime.Nano())
}
