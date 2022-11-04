package clusters

import (
	"fmt"
	"testing"
)

func Test_cluster(t *testing.T) {
	fmt.Println(Cluster([][]int{{1, 2}, {3, 4}, {1, 3}}))
}
