package main

import (
	"fmt"
)

//スライスの操作
func Slice() {
	//長さ・容量の取得
	src := []int{1, 2, 3, 4}
	fmt.Println(src, len(src), cap(src))

	src = append(src, 5)
	fmt.Println(src, len(src), cap(src))

	//値の初期化
	sliceMake := make([]int, 2, 3) //組み込み関数のmakeを用いて長さよ容量を取得
	fmt.Println(sliceMake)

	sliceIndex := []int{2: 1, 5: 5, 7: 13} //インデックスと値を指定する．指定されなかった値はゼロ値で初期化
	fmt.Println(sliceIndex)

	//Slice Trick
	cd := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(cd))
	copy(dst, cd) //スライスの複製
	fmt.Println(dst, len(dst), cap(dst))

	src1, src2 := []int{1, 2}, []int{3, 4, 5}
	cs := append(src1, src2...) //スライス同士の連結
	fmt.Println(cs, len(cs), cap(cs))

	src = []int{1, 2, 3, 4, 5}
	i := 2
	dst = append(src[:i], src[i+1:]...) //3番目の要素の削除
	fmt.Println(dst)

	src = []int{1, 2, 3, 4, 5}
	dst = src[:i+copy(src[i:], src[i+1:])] //appendの代わりにcopyを利用する場合
	fmt.Println(dst)

	src = []int{1, 2, 3, 4, 5}
	//スライスを逆順に並べる
	//方法1
	for i := len(src)/2 - 1; i >= 0; i-- {
		opp := len(src) - 1 - i
		src[i], src[opp] = src[opp], src[i]
	}
	fmt.Println(src)
	//方法2
	for left, right := 0, len(src)-1; left < right; left, right = left+1, right-1 {
		src[left], src[right] = src[right], src[left]
	}
	fmt.Println(src)

	//スライスの要素を偶数のみでフィルタリング
	src = []int{1, 2, 3, 4, 5}
	dst = src[:0]
	for _, v := range src {
		if even(v) {
			dst = append(dst, v)
		}
	}
	fmt.Println(dst)

	//srcをガベージコレクションに回収させることができる
	for i := len(dst); i < len(src); i++ {
		src[i] = 0
	}
	fmt.Println(src)

	//スライスを任意の要素数に分割する
	src = []int{1, 2, 3, 4, 5}
	size := 2
	div := make([][]int, 0, (len(src)+size-1)/size)
	for size < len(src) {
		src, div = src[size:], append(div, src[0:size:size])
	}
	div = append(div, src)
	fmt.Println(div)

}

//引数が偶数かどうかを判定
func even(n int) bool {
	return n%2 == 0
}
