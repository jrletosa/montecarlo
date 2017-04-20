package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const NumShots int = 1000

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	inside := 0.0
	for i := 1; i <= NumShots; i++ {
		x := r.Float64()
		y := r.Float64()
		d := math.Hypot(x, y)
		if d <= 1 {
			inside++
		}
	}

	fmt.Println("Pi: ", 4.0*(inside/float64(NumShots)))
}
