package main

import (
	"fmt"
	"sync"
)

func main() {
	Set1()
}

type Set struct {
	m map[int]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[int]bool{},
	}
}

func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item int) {
	s.Lock()
	s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	fmt.Println("ok:", ok)
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// Set
// 此处只是使用了 int 作为键，还可以实现用 interface{} 作为键，做成更通用的 Set，另外，这个实现是线程安全的。
func Set1() {
	// Go 语言本身是不提供 set 的，但是我们可以自己实现它
	//初始化
	s := New()
	fmt.Println(s.List()) //[]

	//添加元素
	s.Add(1)
	s.Add(1)
	s.Add(3)
	fmt.Println(s.List()) //[1 3]
	if s.IsEmpty() {
		fmt.Println("0 item1")
	}
	//清空元素
	s.Clear()
	fmt.Println(s.List()) //[]

	//判断是否为空
	if s.IsEmpty() {
		fmt.Println("0 item2")
	}

	s.Add(1)
	s.Add(1)
	s.Add(3)
	s.Add(11)
	s.Add(32)

	//判断指定值是否存在
	if s.Has(11) {
		fmt.Println("11 does exist")
	} else {
		fmt.Println("11 doesn't exist")
	}

	fmt.Println(s)
	fmt.Println(s.List())
	fmt.Println(s.Len())

	s.Remove(11)
	s.Remove(3)
	fmt.Println(s.List())
	fmt.Println(s.Len())
}
