package nodeStatus

import (
	"encoding/json"
	"fmt"
	"go-consistent-hashing/controllers"
	"io/ioutil"
	"log"
	"net/http"
)

func GetOneNodeStatus(nodeName string, port string) controllers.NodeStatusStruct {
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

	data := controllers.NodeStatusStruct{}
	json.Unmarshal([]byte(string(body)), &data)
	return data
}

func UpdateNodesStatusMap() {
	theMap := controllers.NodesStatus

	for _, val := range theMap {
		// TODO if no response for timeout set the status as TERMINATED
		data := GetOneNodeStatus(val.NodeName, val.Port)
		// map is updated here
		controllers.NodesStatus[data.Name] = data
	}
	fmt.Println(theMap)
}
