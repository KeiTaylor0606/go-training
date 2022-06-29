package main

import (
	"fmt"
	"testing"
)

func TestAdvancedTextConversion(t *testing.T) {
	EasyTextConversion()
	fmt.Println()
	CharacterCodeAndConversionBetweenHalfWidthAndFullWidth()
	fmt.Println()
	UnicodeAndConversionPerCodePoint()
}
