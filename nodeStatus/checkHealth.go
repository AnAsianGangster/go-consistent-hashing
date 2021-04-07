package nodeStatus

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetOneNodeStatus(port string) {
	resp, err := http.Get("http://localhost:" + port + "/node-health")
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
