package simple

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	s0 := GetInstance()
	s1 := GetInstance()
	if s0 != s1 {
		t.Errorf("s0 not equal s1")
	} else {
		t.Logf("got equal instance")
	}
}

func TestGetInstanceCurrency(t *testing.T) {
	var wg sync.WaitGroup
	instances := make([]*Singleton, 100)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(index int) {
			defer wg.Done()
			instances[index] = GetInstance()
		}(i)
	}

	wg.Wait()

	for i := 1; i < len(instances); i++ {
		if instances[0] != instances[i] {
			t.Errorf("Except got same instance,but got different instance at %d", i)
		}
	}
}
