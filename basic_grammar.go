package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("My favorite number is", rand.Intn(10))
	}
	j := 0
	for j < 1 {
		rnd := rand.Intn(10)
		if rnd < 5 || rnd == 10 {
			fmt.Println("While loop", rnd)
			j += 1
		}
	}
	item := [3]int{4, 5, 6}
	for i := range item {
		fmt.Println(item[i], i)
	}

	for i := range item {
		fmt.Println(item[i], i)
	}
}

func rand_methods() {
	// int型の 0 ~ n - 1 の値を生成
	rnd_num := rand.Intn(10)
	fmt.Println("Random Number is", rnd_num)
}
