package wiggle_max_length

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/go-playground/assert/v2"
)

func Test_wiggleMaxLength(t *testing.T) {
	c := struct{ name string }{name: "zhangsan"}
	d := struct{ name string }{name: "zhangsan"}

	fmt.Println(uintptr(unsafe.Pointer(&c)), uintptr(unsafe.Pointer(&d)))
	assert.Equal(t, 3, wiggleMaxLength([]int{1, 3, 2}))
}
