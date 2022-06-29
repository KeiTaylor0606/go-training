package advancedgrammer

import "testing"

func TestHello(t *testing.T) {
	go hello("hello", 50)
	hello("bye!", 100)
}
