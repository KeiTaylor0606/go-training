package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/rivo/uniseg"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
)

func EasyTextConversion() {
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

func CharacterCodeAndConversionBetweenHalfWidthAndFullWidth() {
	/*
		transform.Transformerインターフェース
	*/

	//transoform.NewReader関数の使用例
	r := strings.NewReader("Hello, World")
	tr := transform.NewReader(r, transform.Nop) //transform.Nop変数は何も変換を行わないtransform.Transformer
	_, err := io.Copy(os.Stdout, tr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	//transoform.NewWriter関数の使用例
	r = strings.NewReader("Hello, World")
	tw := transform.NewWriter(os.Stdout, transform.Nop)
	_, err = io.Copy(tw, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	/*
		文字コードの変換
	*/

	/*
		半角と全角の変換
	*/

	//East Asian Width特性を取得
	rs := []rune{'５', 'ァ', 'ア', 'A', 'α'}
	fmt.Println("rune\tWide\tNarrow\tFolded\tKind")
	fmt.Println("----------------------------------------------")
	for _, r := range rs {
		p := width.LookupRune(r)
		w, n, f, k := p.Wide(), p.Narrow(), p.Folded(), p.Kind()
		fmt.Printf("%2c\t%2c\t%3c\t%3c\t%s\n", r, w, n, f, k)
	}

	//通常半角で表す文字は半角に，通常全角で表す文字は全角に変換
	for _, r := range width.Fold.String("５ァアAα") {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}

	//全角文字を半角文字に変換
	for _, r := range width.Narrow.String("５ァアAα") {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}

	//半角文字を全角文字に変換
	for _, r := range width.Widen.String("５ァアAα") {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}

	/*
		Unicodeの正規化
	*/
}

func UnicodeAndConversionPerCodePoint() {
	/*
		UnicodeとUTF-8
	*/
	for i, r := range "世界" {
		fmt.Printf("%d: %c", i, r)
		fmt.Println()
	}

	//UnicodeコードポイントをUTF-8にエンコード
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, '世')
	fmt.Printf("%v %q %d\n", buf, string(buf), n)

	//UTF-8をUnicodeコードポイントにデコード
	b := []byte("世界")
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%q:%v \n", r, size)
		b = b[size:]
	}

	//文字列"Cafe\u0301"を書記素クラスタに分解
	gr := uniseg.NewGraphemes("Cafe\u0301")
	for gr.Next() {
		fmt.Printf("%s %x \n", gr.Str(), gr.Runes())
	}
}

/*
Shift_JISのCSVファイルをUTF-8に変換して読み込む
*/
func printCSV(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	//Shift_JISとして読み込む
	dec := japanese.ShiftJIS.NewDecoder()
	cr := csv.NewReader(dec.Reader(f))
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(rec) //UTF-8に変換されているので表示しても文字化けしない
	}
	return nil
}

/*
文字コード変換と全角半角変換を行う関数
*/
func foldShiftJISFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	//Shift_JISからUTF-8に変換してから全角英数などは半角に，半角カナなどは全角にする
	dec := japanese.ShiftJIS.NewDecoder()
	t := transform.Chain(dec, width.Fold)
	r := transform.NewReader(f, t)

	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}
