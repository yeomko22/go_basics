package ch2_sync

import (
	"fmt"
	"sync"
	"testing"
)

func poolBasic() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}

func TestPool(t *testing.T) {
	poolBasic()
}
