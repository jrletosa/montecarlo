package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Read number of shots from command line
	numShots := flag.Int("shots", 100, "number of random shots")
	flag.Parse()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	inside := 0.0
	for i := 1; i <= *numShots; i++ {
		x := r.Float64()
		y := r.Float64()
		d := math.Hypot(x, y)
		if d <= 1 {
			inside++
		}
	}

	fmt.Println("Pi: ", 4.0*(inside/float64(*numShots)))
}
