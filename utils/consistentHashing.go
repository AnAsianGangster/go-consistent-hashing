package utils

import "fmt"

const (
	MAX_VAL = ^uint64(0)
)

func GetNodeLocation(numOfNodes int, key string) int {
	nodeCircleArray := make([]uint64, numOfNodes)

	for i := 0; i < numOfNodes; i++ {
		// 64* bit / number of nodes * (node id)
		nodeCircleArray[i] = uint64(MAX_VAL / uint64(numOfNodes) * uint64(i))
	}

	// When the client ask the node
	hash := GetHashValue(key)
	fmt.Println(hash)
	var idx int
	for i := numOfNodes - 1; i >= 0; i-- {
		// 64* bit / number of nodes * (node id)
		if hash > nodeCircleArray[i] {
			idx = i
			break
		}
	}
	// Ask if server idx is alive
	// Return idx value to the client
	fmt.Printf("\033[31mKey %v is stored in Node: %v\033[0m\n", key, idx)
	return idx
}
