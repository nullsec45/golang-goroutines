package go_routines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for time := range ticker.C {
		fmt.Println(time)
	}
}
