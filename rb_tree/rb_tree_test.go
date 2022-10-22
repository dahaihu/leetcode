package rb_tree

import (
	"fmt"
	"testing"
)

func Test_rbTree(t *testing.T) {
	rbTree := &RBTree{Compare: func(a interface{}, b interface{}) int {
		aVal, bVal := a.(int), b.(int)
		if aVal > bVal {
			return 1
		} else if aVal < bVal {
			return -1
		} else {
			return 0
		}
	}}
	for i := 0; i < 10; i++ {
		rbTree.Set(i, i)
		fmt.Println(rbTree)
		fmt.Println(rbTree.size)
	}
	fmt.Println(rbTree)
	rbTree.Remove(9)
	fmt.Println(rbTree)
	rbTree.Remove(4)
	fmt.Println(rbTree)
	rbTree.Remove(7)
	fmt.Println(rbTree)
	rbTree.Remove(3)
	fmt.Println(rbTree)
	fmt.Println(rbTree.size)
}

func Test_colorPrint(t *testing.T) {
	colorReset := "\033[0m"

	//colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"

	fmt.Printf("\033[31m%s", "test")
	fmt.Println(string(colorGreen), "test")
	fmt.Println(string(colorYellow), "test")
	fmt.Println(string(colorBlue), "test")
	fmt.Println(string(colorPurple), "test")
	fmt.Println(string(colorWhite), "test")
	fmt.Println(string(colorCyan), "test", string(colorReset))
	fmt.Println("next")
	fmt.Println(len("\u001B[31m"))
}

func Test_color(t *testing.T) {
}
