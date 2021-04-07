package nodeStatus

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ChangeNodeStatus(port string) {
	requestBody, err := json.Marshal(map[string]string{
		"status": "TERMINATING",
	})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:"+port+"/node-status", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))

}
