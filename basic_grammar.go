package main

import (
	"fmt"
	"math/rand"
)

func main() {

	rand_methods()

	for_templates()

}

func rand_methods() {
	// int型の 0 ~ n - 1 の値を生成
	rnd_num := rand.Intn(10)
	fmt.Println("Random Number is", rnd_num)
}

func for_templates() {
	/* 馴染み深い for ------------------------------------------ */
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	/* While的な書き方 ----------------------------------------- */
	j := 0
	for j < 3 {
		rnd := rand.Intn(10)
		if rnd < 5 || rnd == 10 {
			fmt.Println("While loop", rnd)
			j += 1
		}
	}

	/* 配列を見る --------------------------------------------- */
	item := [3]int{4, 5, 6}
	for i := range item {
		fmt.Println(item[i], i)
	}
}
