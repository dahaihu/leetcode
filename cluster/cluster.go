package clusters

func ancestor(parents map[int]int, target int) int {
	for {
		parent := parents[target]
		if parent == target {
			return parent
		}
		target = parent
	}
}

func cluster(nums [][]int) [][]int {
	clusters := make(map[int][]int)
	parents := make(map[int]int)
	for _, num := range nums {
		for _, node := range num {
			if _, ok := parents[node]; !ok {
				parents[node] = node
			}
			if _, ok := clusters[node]; !ok {
				clusters[node] = []int{node}
			}
		}
	}

	for _, num := range nums {
		child, parent := num[0], num[1]
		childAncestor := ancestor(parents, child)
		parentAncestor := ancestor(parents, parent)
		if childAncestor != parentAncestor {
			clusters[parentAncestor] = append(
				clusters[parentAncestor],
				clusters[childAncestor]...)
			delete(clusters, childAncestor)
			parents[childAncestor] = parentAncestor
		}
	}
	out := make([][]int, 0)
	for _, cluster := range clusters {
		out = append(out, cluster)
	}
	return out
}
