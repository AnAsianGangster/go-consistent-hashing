package nodeStatus

import (
	"encoding/json"
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
	// log.Println(string(body)) // TODO return this as json
	json.Unmarshal([]byte(string(body)), &data)
	// fmt.Println(data)
	return data
}
