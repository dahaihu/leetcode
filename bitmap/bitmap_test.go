package bitmap

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kelindar/bitmap"
)

func Test_bitMap(t *testing.T) {
	var books bitmap.Bitmap

	books.Set(3)                 // Set the 3rd bit to '1'
	hasBook := books.Contains(3) // Returns 'true'
	fmt.Println(hasBook)
	books.Remove(3) // Set the 3rd bit to '0'
	reflect.TypeOf(books)
}
