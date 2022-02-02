package main

import (
	"fmt"
	"time"
)

func Type() {
	//構造体の定義
	var empty struct{} //フィールドを持たない構造体
	fmt.Println(empty)

	var point struct { //フィールドを三つ持つ構造体
		ID   string
		x, y int
	}
	fmt.Println(point)

	//配列の定義
	var array [5]int //ゼロ値で初期化
	fmt.Println(array)

	arrayLiteral := [5]int{1, 2, 3, 4, 5} //5つの要素を持つ配列を定義
	fmt.Println(arrayLiteral)

	arrayInference := [...]int{1, 2, 3, 4, 5} //要素数から配列数を推論，この場合は5
	fmt.Println(arrayInference)

	arrayIndex := [...]int{2: 1, 5: 5, 7: 13} //配列のインデックスと値を指定．インデックスの指定がない箇所はゼロ
	fmt.Println(arrayIndex)

	//スライスの定義
	var slice []int //ゼロ値で初期化
	fmt.Println(slice)

	sliceLiteral := []int{1, 2, 3, 4, 5} //5つの要素を持つスライスを定義
	fmt.Println(sliceLiteral)

	//マップの定義
	var m map[string]int //ゼロ値で初期化
	fmt.Println(m)

	mapLiteral := map[string]int{ //2つの要素を持つマップの定義
		"John":    42,
		"Richard": 33,
	}
	fmt.Println(mapLiteral)

	//ユーザ定義
	type MyDuration time.Duration //新しい型MyDurationをtime.Durationを基底として定義
	d := MyDuration(100)

	fmt.Printf("%T", d) //%Tを使うことで値の型情報を出力できる

	td := time.Duration(d) //MyDuration型で基底型として定義しているtime.Durationへのキャスト

	md := 100 * d //型の定義がされていない定数（100）に対して明示的なキャストなしでの演算

	fmt.Printf("td: %T, md %T", td, md)
}
