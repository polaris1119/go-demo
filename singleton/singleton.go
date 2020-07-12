// 懒汉式单例模式
package singleton

import (
	"sync"
)

type singleton struct {
	count int
}

var (
	instance *singleton
	mutex    sync.Mutex
)

// New 普通版本
func New() *singleton {
	mutex.Lock()
	if instance == nil {
		instance = new(singleton)
	}
	mutex.Unlock()

	return instance
}

// New2 双重检查版本
func New2() *singleton {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = new(singleton)
		}
		mutex.Unlock()
	}

	return instance
}

func (s *singleton) Add() int {
	s.count++
	return s.count
}

var once sync.Once

func New3() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}
