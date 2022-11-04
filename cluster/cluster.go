package clusters

func findAncestor(parents map[int]int, target int) int {
	for parents[target] != target {
		target = parents[target]
	}
	return target
}

func Cluster(nums [][]int) [][]int {
	// init
	clusters := make(map[int][]int)
	parents := make(map[int]int)
	for _, num := range nums {
		for _, ele := range num {
			if _, ok := clusters[ele]; !ok {
				clusters[ele] = []int{ele}
			}
			if _, ok := parents[ele]; !ok {
				parents[ele] = ele
			}
		}
	}

	// union
	for _, num := range nums {
		child, parent := num[0], num[1]
		childAncestor := findAncestor(parents, child)
		parentAncestor := findAncestor(parents, parent)
		clusters[parentAncestor] = append(clusters[parentAncestor], clusters[childAncestor]...)
		delete(clusters, childAncestor)
		parents[childAncestor] = parentAncestor
	}

	// finish
	var result [][]int
	for _, c := range clusters {
		result = append(result, c)
	}
	return result
}
