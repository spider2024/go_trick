package simple

import "sync"

type Singleton struct {
}

var instance *Singleton

var m sync.Mutex

func GetInstance() *Singleton {
	if instance == nil {
		m.Lock()
		if instance == nil {
			instance = &Singleton{}
		}
		defer m.Unlock()
	}
	return instance
}
