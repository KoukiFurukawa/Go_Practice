package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	rand_methods()        // ランダム
	for_templates()       // for 文
	func_arguments(1, 2)  // 関数の引数
	parse_samples()       // int->string string->int
	named_return_value(1) // 戻り値の指定
	variable_modifier()   // 変数の設定
	switch_sample()       // ランタイムを確認する
	defer_sample()        // 終わってから動作
	struct_sample()       // 構造体テスト
	Array_sample()        // 配列
	slice_operation()     // slice操作
	map_sample()          // 辞書型
	compute(swap)         // 関数を引数に
	add(10, 11)           // 別ファイル呼び出し
	createUserInfo(
		[]uint8{17, 18, 19, 59, 60, 61}) // user作成
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
	for i, v := range item {
		fmt.Println(item[i], i, v)
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

// Swap する
func swap(x int, y int) (int, int) {
	return y, x
}

// 戻り値の型と名前の指定を行える
// 短い関数でのみ行うべき
func named_return_value(sum int) (x, y int) {
	x = sum * 4
	y = sum / 4
	return
}

func variable_modifier() {
	// var 宣言で変数
	var i, j, k int // 0が代入される
	var b bool      // false
	var s string    // ""
	i = 3
	j = 1
	fmt.Println(i, j, k)
	fmt.Printf("Type: %T\n Value:%v\n", b, b)
	fmt.Printf("Type: %T\n Value:%v\n", s, s)

	// var宣言で初期化子が与えられてる時、型の指定は省略できる
	var char, number = "string", 0
	fmt.Println(char, number)

	// 関数内では、 := で同じことができる。
	message := "ok"
	fmt.Println(message)

	// const で定数
	const v string = "Value"
	const num int = 3
	const flag bool = true
}

func while() {
	for {
		// inf loop
	}
}

func conditional_branch() {
	a := 0
	b := 1
	if a == b {
		fmt.Println("a is true")
	} else if a == 0 {
		return
	} else {
		return
	}

}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func switch_sample() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func defer_sample() {
	hello()
	// defer は **スタック** される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func hello() {
	defer fmt.Println("Hello")
	fmt.Print("World")
}

type Human struct {
	age  int
	name string
}

func struct_sample() {
	var human = Human{age: 18, name: "john"}
	fmt.Println(human)
}

func Array_sample() {
	var a [2]string
	a[0] = "fizz"
	a[1] = "buzz"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// スライス
	var slice []int = primes[1:4]
	make_slice := make([]int, 3)
	fmt.Println(primes, slice)
	fmt.Println(len(primes), cap(make_slice))

	// スライス色々
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func slice_operation() {
	// append
	s := make([]int, 10)
	s = append(s, 1)
	fmt.Println(s)
	s = append(s, 2, 3, 4, 5)
	fmt.Println(s)

	// slice_object
	so := make([][]uint8, 10)
	so[0] = make([]uint8, 10)
	fmt.Println(cap(so[0]), so)
}

func map_sample() {
	m := map[string]int{"hey": 4}
	println(m["hey"])
	// 初期化しておく
	mp := make(map[string]Human)
	mp["super"] = Human{age: 44, name: "man"}
	fmt.Println(mp["super"])

	// hey-4 削除
	delete(m, "hey")

	// hey-4 してない？
	elem, ok := m["hey"]
	fmt.Println("the value:", elem, "Present?", ok)
}

func compute(fn func(int, int) (int, int)) int {
	var r, l int
	r = 1
	l = 10
	r, l = fn(r, l)
	fmt.Println(r, l)
	return r + l
}
