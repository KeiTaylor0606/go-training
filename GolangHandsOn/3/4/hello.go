package advancedgrammer

import (
	"fmt"
	"time"
)

func hello(s string, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("<%d %s>\n", i, s)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}
