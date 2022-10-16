package full_justify

import (
	"fmt"
	"strings"
	"testing"
)

func Test_fullLine(t *testing.T) {
	fmt.Println(strings.Join(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16), "\n"))
}

func Test_mark(t *testing.T) {
}
