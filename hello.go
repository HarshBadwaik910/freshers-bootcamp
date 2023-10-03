package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/maps"
)

func main() {
	a := []int{1,2,3,4,5}
	b := []int{1,2,3,4,5}
	fmt.Println(slices.Equal(a,b))
	c := map[string]int{"foo":1,"bar":2}
	d := map[string]int{"foo":1,"bar":3}
	fmt.Println(maps.Equal(c,d))
}