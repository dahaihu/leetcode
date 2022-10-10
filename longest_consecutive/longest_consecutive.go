package longest_consecutive

var placeholder = struct{}{}

type Set map[int]struct{}

func NewSet(elements ...int) Set {
	s := make(Set)
	for _, element := range elements {
		s[element] = placeholder
	}
	return s
}

func (s Set) Add(ele int) {
	s[ele] = placeholder
}

func (s Set) Remove(ele int) {
	delete(s, ele)
}

func (s Set) Merge(other Set) {
	for ele := range other {
		s[ele] = placeholder
	}
}

func (s Set) Len() int {
	return len(s)
}

func rootAncestor(mark map[int]int, ele int) int {
	for {
		parent := mark[ele]
		if ele == parent {
			return parent
		}
		ele = parent
	}
}

func longestConsecutive(nums []int) int {
	ancestors := make(map[int]int)
	children := make(map[int]Set)
	var maxLength int
	for _, num := range nums {
		var root int
		{
			_, leftOk := ancestors[num-1]
			if leftOk {
				root = rootAncestor(ancestors, num-1)
			} else {
				root = num
			}
		}
		rootChildren, ok := children[root]
		if !ok {
			rootChildren = NewSet(root)
			children[root] = rootChildren
		}
		rootChildren.Add(num)
		ancestors[num] = root
		{
			rightChildren, ok := children[num+1]
			if ok {
				ancestors[num+1] = num
				rootChildren.Merge(rightChildren)
			}
		}
		if length := rootChildren.Len(); length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func longestConsecutive1(nums []int) int {
	mark := make(map[int]struct{})
	for _, num := range nums {
		mark[num] = placeholder
	}
	var maxLength int
	for _, num := range nums {
		_, ok := mark[num-1]
		if ok {
			continue
		}
		curLength := 1
		for {
			num += 1
			_, ok = mark[num]
			if !ok {
				break
			}
			curLength += 1
		}
		if curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}
