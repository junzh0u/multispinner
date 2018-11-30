package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/junzh0u/multispinner"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	group := multispinner.NewSpinnerGroup(10)
	group.Start()
	for i := 0; i < 10; i++ {
		go func(idx int) {
			time.Sleep(time.Duration(rand.Int()%5000) * time.Millisecond)
			dice := rand.Int()%6 + 1
			message := fmt.Sprintf("Rolled a %d!", dice)
			if dice <= 3 {
				group.At(idx).Error(message)
			} else {
				group.At(idx).Success(message)
			}
		}(i)
	}
	group.Wait()
}
