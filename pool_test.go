package go_routines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	//var pool sync.Pool
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Rama")
	pool.Put("Fajar")
	pool.Put("Fadhillah")

	for i := 0; i < 3; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
		//data := pool.Get()
		//fmt.Println(data)
		//pool.Put(data)
	}
	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
