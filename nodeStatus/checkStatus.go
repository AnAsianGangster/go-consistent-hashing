package nodeStatus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetOneNodeStatus(nodeName string, port string) NodeStatusStruct {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Get("http://" + nodeName + ":" + port + "/node-health")
	if err != nil {
		log.Println(err)
		return NodeStatusStruct{}
	}
	// TODO if no response for timeout set the status as TERMINATED
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	data := NodeStatusStruct{}
	json.Unmarshal([]byte(string(body)), &data)
	return data
}

func UpdateNodesStatusMap() {
	theMap := NodesStatus

	for _, val := range theMap {
		data := GetOneNodeStatus(val.NodeName, val.Port)
		// if it return empty do not add the empty struct to the map
		emptyNodeStruct := NodeStatusStruct{}
		if data != emptyNodeStruct {
			// map is updated here
			NodesStatus[data.NodeName] = data
		} else {
			// remove the node from the map
			delete(theMap, val.NodeName)
		}
	}
	fmt.Println(theMap)
}
