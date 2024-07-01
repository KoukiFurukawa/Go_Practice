package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
)

func main() {

	rand_methods()

	for_templates()

	func_arguments(1, 2)

	parse_samples()

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

/* 引数あり --------------------------------------------------- */
func func_arguments(x int, y int) int {
	return x + y
}
func func_arguments2(x, y int, z string) string {
	xy := x * y
	completion_string := z + fmt.Sprint(xy)
	return completion_string
}

func parse_samples() {
	number := 1e9 + 7 // 1000000007 素数である
	char := "57"

	/* string に変換 */
	num_to_string := fmt.Sprint(number)
	fmt.Println(num_to_string, "%T", num_to_string)

	/* int に変換 */
	string_to_num, e := strconv.ParseInt(char, 10, 64)
	if e == nil {
		fmt.Println(string_to_num, reflect.TypeOf(string_to_num))
	}
}
