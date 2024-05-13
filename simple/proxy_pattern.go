package simple

import (
	"fmt"
	"sync"
)

type Seller interface {
	sell(name string)
}

type Station struct {
	stock int
	lock  sync.Mutex
}

func (station *Station) sell() bool {
	station.lock.Lock()
	defer station.lock.Unlock()
	if station.stock > 0 {
		station.stock--
		return true
	} else {
		return false
	}
}

type StationProxy struct {
	station *Station
	lock    sync.Mutex
}

func (proxy *StationProxy) sell() bool {
	proxy.lock.Lock()
	defer proxy.lock.Unlock()
	if proxy.station.stock > 0 {
		proxy.station.stock--
		return true
	} else {
		fmt.Println("over...")
		return false
	}
}
