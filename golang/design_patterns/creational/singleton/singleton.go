package singleton

import "sync"

type Singleton struct {
	name string
}

var instance *Singleton

var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func (s *Singleton) SetName(name string) {
	s.name = name
}

func (s *Singleton) GetName() string {
	return s.name
}
