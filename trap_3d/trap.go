package trap_3d

import "container/heap"

type queue struct {
	elements [][]int
}

func (q *queue) Less(i, j int) bool {
	return q.elements[i][2] <= q.elements[j][2]
}

func (q *queue) Len() int {
	return len(q.elements)
}

func (q *queue) Swap(i, j int) {
	q.elements[i], q.elements[j] = q.elements[j], q.elements[i]
}

func (q *queue) Push(ele interface{}) {
	q.elements = append(q.elements, ele.([]int))
}

func (q *queue) Pop() interface{} {
	last := q.elements[q.Len()-1]
	q.elements = q.elements[:q.Len()-1]
	return last
}

func trapRainWater(heightMap [][]int) int {
	xLen, yLen := len(heightMap), len(heightMap[0])
	if xLen <= 2 || yLen <= 2 {
		return 0
	}
	var q queue
	visited := make(map[[2]int]bool)
	for i := 0; i < xLen; i++ {
		for j := 0; j < yLen; j++ {
			if i == 0 || j == 0 || i == xLen-1 || j == yLen-1 {
				heap.Push(&q, []int{i, j})
				visited[[2]int{i, j}] = true
			}
		}
	}
	var waters int
	nexts := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for q.Len() != 0 {
		min := heap.Pop(&q).([]int)
		x, y, h := min[0], min[1], heightMap[min[0]][min[1]]
		for _, next := range nexts {
			nextX, nextY := x+next[0], y+next[1]
			if nextX >= 0 && nextX < xLen &&
				nextY >= 0 && nextY < yLen &&
				!visited[[2]int{nextX, nextY}] {
				nextH := heightMap[nextX][nextY]
				if cur := h - nextH; cur > 0 {
					waters += cur
					nextH = h
				}
				heap.Push(&q, []int{nextX, nextY, nextH})
				visited[[2]int{nextX, nextY}] = true
			}
		}
	}
	return waters
}
