package urand

import (
	"math/rand"
	"time"
)

//------------------------------public method------------------------------

type URand struct {
	list  []int64
	index int
}

func New(size int) *URand {
	instance := &URand{
		list:  make([]int64, size),
		index: 0,
	}

	rand.Seed(time.Now().UnixNano())
	instance._init()
	return instance
}

func (self *URand) Get() int64 {
	if self.index >= len(self.list) {
		self.index = 0
	}

	list := self.list
	index := self.index
	targetIndex := index + rand.Intn(len(self.list)-index)

	v := list[targetIndex]
	list[index], list[targetIndex] = v, list[index]

	self.index++

	return v
}

//------------------------------private method------------------------------
func (self *URand) _init() {
	for i, _ := range self.list {
		self.list[i] = int64(i)
	}
}
