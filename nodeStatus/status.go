package nodeStatus

// TODO change to enums & move to a private folder
var NodesStatus = make(map[string]NodeStatusStruct)

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
