package nodeStatus

import (
	"fmt"
)

// TODO change to enums?
// map of the nodes
var NodesStatus = make(map[string]NodeStatusStruct)
// mapping of the node with int index
var NodeIdxNameMap = make(map[int]string)

// ANCHOR I use Content-Type: application/json. Might need to change to postform
type NodeStatusStruct struct {
	Name     string `json:"name"`
	NodeName string `json:"nodeName"`
	Port     string `json:"port"`
	Status   string `json:"status"`
}

func GetNumberOfAliveNodes() int {
	return len(NodesStatus)
}

func UpdateNodesStatusMapToArrayMapping() {
	idx := 0
	for _, val := range NodesStatus {
		NodeIdxNameMap[idx] = val.NodeName
		idx++
	}
	fmt.Println(NodeIdxNameMap)
}

