package nodeStatus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetOneNodeStatus(nodeName string, port string) NodeStatusStruct {
	resp, err := http.Get("http://" + nodeName + ":" + port + "/node-health")
	if err != nil {
		log.Fatal(err)
	}

	// TODO if no response for timeout set the status as TERMINATED
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := NodeStatusStruct{}
	json.Unmarshal([]byte(string(body)), &data)
	return data
}

func UpdateNodesStatusMap() {
	theMap := NodesStatus

	for _, val := range theMap {
		// TODO if no response for timeout set the status as TERMINATED
		data := GetOneNodeStatus(val.NodeName, val.Port)
		// map is updated here
		NodesStatus[data.Name] = data
	}
	fmt.Println(theMap)
}
