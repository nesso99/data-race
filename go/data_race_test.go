package data_race_test

import (
	"sync"
	"testing"
)

func TestDataRace(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	x := 1

	go func() {
		for firstIterator := 0; firstIterator < 500000; firstIterator++ {
			x += 1
		}
		wg.Done()
	}()

	go func() {
		for secondIterator := 0; secondIterator < 500000; secondIterator++ {
			x += 1
		}
		wg.Done()
	}()

	wg.Wait()
	if x == 1000000 {
		t.Errorf("Failed")
	} else {
		t.Logf("x is %d", x)
	}
}
