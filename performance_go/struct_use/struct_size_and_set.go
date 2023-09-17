package main

import (
	"fmt"
	"unsafe"
)

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func main() {
	// 计算空结构体的大小
	fmt.Println(unsafe.Sizeof(struct{}{})) // 0

	s := make(Set)
	s.Add("SteveRocket")
	s.Add("CTO Plus")
	fmt.Println(s.Has("Cramer"))  // false
	fmt.Println(s.Has("SteveRocket"))  // true
}
