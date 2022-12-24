package clusters

func ancestor(parents map[int]int, target int) int {
	p := parents[target]
	for p != target {
		target = p
		p = parents[target]
	}
	return p
}

func cluster(nums [][]int) [][]int {
	clusters := make(map[int][]int)
	parents := make(map[int]int)
	for _, num := range nums {
		for _, ele := range num {
			if _, ok := parents[ele]; !ok {
				parents[ele] = ele
			}
			if _, ok := clusters[ele]; !ok {
				clusters[ele] = []int{ele}
			}
		}
	}
	for _, num := range nums {
		child, parent := num[0], num[1]
		childAncestor := ancestor(parents, child)
		parentAncestor := ancestor(parents, parent)
		clusters[parentAncestor] = append(clusters[parentAncestor],
			clusters[childAncestor]...)
		delete(clusters, childAncestor)
		parents[child] = parent
	}
	var result [][]int
	for _, cluster := range clusters {
		result = append(result, cluster)
	}
	return result
}
