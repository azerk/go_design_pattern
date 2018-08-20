package main

import (
	"sync"
	"fmt"
)

func main() {
	s:=New()
	s["hello"]="world"
	s1:=New()

	pro:=GetProduct()
	pro.name="A"
	pro.count=10
	pro1:=GetProduct()

	fmt.Println(s1["hello"])
	fmt.Println(pro1.name)
}

type singleton map[string]string


type product struct {
	name string
	count int
}

func GetProduct() *product  {
	once1.Do(func() {
		pro=&product{}
	})
	return pro
}

var (
	once,once1 sync.Once
	instance singleton
	pro *product
)
func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}
