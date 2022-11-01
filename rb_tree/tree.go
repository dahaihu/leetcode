package rb_tree

import "strings"

type INode interface {
	IsNil() bool
	String() string
	LeftChild() INode
	RightChild() INode
}

func midIndex(start, end int, leftBox bool) int {
	totalCount := end - start + 1
	count, remainder := totalCount/2, totalCount%2
	if remainder == 1 {
		return start + count
	}
	if leftBox {
		return start + count
	}
	return start + count - 1
}

func BuildTree(node INode) (
	box []string, boxLength, rootStart, rootEnd int,
) {
	if node.IsNil() {
		return nil, 0, 0, 0
	}
	rootStr := node.String()
	rootWidth := len(rootStr)
	gapSize := len(rootStr)

	leftBox, leftBoxLength, leftRootStart, leftRootEnd := BuildTree(node.LeftChild())
	rightBox, rightBoxLength, rightRootStart, rightRootEnd := BuildTree(node.RightChild())
	var (
		line1, line2 []string
	)
	if leftBoxLength > 0 {
		leftRootIndex := midIndex(leftRootStart, leftRootEnd, true)
		line1 = append(line1, strings.Repeat(" ", leftRootIndex+1))
		line1 = append(line1, strings.Repeat("_", leftBoxLength-leftRootIndex))
		line2 = append(line2, strings.Repeat(" ", leftRootIndex)+"/")
		line2 = append(line2, strings.Repeat(" ", leftBoxLength-leftRootIndex))
		rootStart = leftBoxLength + 1
		gapSize++
	} else {
		rootStart = 0
	}
	line1 = append(line1, rootStr)
	line2 = append(line2, strings.Repeat(" ", rootWidth))
	if rightBoxLength > 0 {
		rightRootIndex := midIndex(rightRootStart, rightRootEnd, false)
		line1 = append(line1, strings.Repeat("_", rightRootIndex+1))
		line1 = append(line1, strings.Repeat(" ",
			rightBoxLength-rightRootIndex+1))
		line2 = append(line2, strings.Repeat(" ", rightRootIndex+1)+"\\")
		line2 = append(line2, strings.Repeat(" ",
			rightBoxLength-rightRootIndex))
		gapSize++
	}
	rootEnd = rootStart + rootWidth - 1
	gapStr := strings.Repeat(" ", gapSize)
	newBox := []string{strings.Join(line1, ""), strings.Join(line2, "")}
	childHeight := len(leftBox)
	if rightBoxHeight := len(rightBox); rightBoxHeight > childHeight {
		childHeight = rightBoxHeight
	}
	for i := 0; i < childHeight; i++ {
		var lline, rline string
		if i < len(leftBox) {
			lline = leftBox[i]
		} else {
			lline = strings.Repeat(" ", leftBoxLength)
		}
		if i < len(rightBox) {
			rline = rightBox[i]
		} else {
			rline = strings.Repeat(" ", rightBoxLength)
		}
		newBox = append(newBox, lline+gapStr+rline)
	}
	return newBox, len(newBox[0]), rootStart, rootEnd
}
