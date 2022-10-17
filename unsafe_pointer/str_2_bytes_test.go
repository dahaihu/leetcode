package unsafe_pointer

import (
	"fmt"
	"testing"
)

func Test_array(t *testing.T) {
	a := [3]int{1, 2, 3}
	ad := &a
	fmt.Println(ad[0], ad[1])
}

func Test_str2bytes(t *testing.T) {
	s := "Hello world"
	for _, ele := range str2bytes(s) {
		fmt.Printf("%c", ele)
	}
}

func Test_bytes2str(t *testing.T) {
	bs := []byte("Hello world")
	fmt.Println(bytes2str(bs))
}
