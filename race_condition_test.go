package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)
func TestRaceCondition(t *testing.T) {
	x := 0
	
	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
				// ada dua/lebih goroutine mengakses data x bersamaan saat melakukan counter (karena parallel programming)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)
}