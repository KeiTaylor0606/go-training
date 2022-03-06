package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"
)

func AdvancedTextConversion() {
	/*
		stringsパッケージ
	*/

	//文字列の置換
	s := strings.Replace("郷に入っては郷に従え", "郷", "Go", 1)
	fmt.Println(s)

	//文字列の置換（第二引数で指定した文字列全てを置換する）
	s = strings.Replace("郷に入っては郷に従え", "郷", "Go", -1)
	fmt.Println(s)

	//複数パターンの置換
	r := strings.NewReplacer("郷", "Go", "入れば", "入っては")
	s = r.Replace("郷に入れば郷に従え")
	fmt.Println(s)

	//小文字の英字を大文字に変換
	toUpper := func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return r - 'a' + 'A'
		}
		return r
	}
	s = strings.Map(toUpper, "Hello, World")
	fmt.Println(s)

	//小文字を大文字に変換（unicode.ToUpper関数）
	s = strings.Map(unicode.ToUpper, "Hello, World")
	fmt.Println(s)

	/*
		bytesパッケージ
	*/

	//[]byte型の値の変換
	b := bytes.ReplaceAll([]byte{0x0A, 0x0B, 0x0C}, []byte{0x0B}, []byte{0xFF})
	fmt.Printf("% X\n", b)

	/*
		regexpパッケージ
	*/

	//正規表現を使った置換
	re, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		log.Fatal(err)
	}
	s = re.ReplaceAllString("1986年01月12日", "${2}/${3} ${1}")
	fmt.Println(s)

	//正規表現を使った置換（キャプチャした文字列を展開しない）
	re, err = regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		log.Fatal(err)
	}
	s = re.ReplaceAllLiteralString("1986年01月12日", "${2}/${3} ${1}")
	fmt.Println(s)

	//正規表現にマッチした文字列を，関数で指定したルールで変換
	re, err = regexp.Compile(`(^|_)[a-zA-Z]`)
	if err != nil {
		log.Fatal(err)
	}
	s = re.ReplaceAllStringFunc("hello_world", func(s string) string {
		return strings.ToUpper(strings.TrimLeft(s, "_"))
	})
	fmt.Println(s)
}
