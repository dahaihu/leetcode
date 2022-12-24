package clusters

import (
	"fmt"
	"testing"
)

func Test_cluster(t *testing.T) {
	fmt.Println(cluster([][]int{{1, 2}, {3, 4}}))
}
