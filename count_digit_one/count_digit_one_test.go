package countdigitone

import (
	"fmt"
	"testing"
)

func Test_countDigitOne(t *testing.T) {
	fmt.Println(countDigitOne(11))
	fmt.Println(countDigitOne(808182) == 503522)
}
