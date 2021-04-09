package main

import (
	"fmt"
	"go-consistent-hashing/controllers"
	"go-consistent-hashing/nodeStatus"
	"go-consistent-hashing/routers"
	"go-consistent-hashing/utils"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// ping for nodes
	go func() {
		for {
			theMap := controllers.NodesStatus
			// for key, val := range theMap {
			for _, val := range theMap {
				// fmt.Println("Key:", key, "=>", "val:", val)
				data := nodeStatus.GetOneNodeStatus(val.NodeName, val.Port)
				controllers.NodesStatus[data.Name] = data
			}
			fmt.Println(theMap)
			// fmt.Println(time.Now().UTC())
			// sleep for one second
			time.Sleep(time.Second * 1)
		}
	}()

	// TODO Listen at some port
	// TODO mapping for node name with index in the node_value_mapping
	// number_of_nodes := len(controllers.NodesStatus)
	number_of_nodes := 10
	node_value_mapping := make([]uint64, number_of_nodes)
	MAX_VAL := ^uint64(0)

	for i := 0; i < number_of_nodes; i++ {
		// 64* bit / number of nodes * (node id)
		node_value_mapping[i] = uint64(MAX_VAL / uint64(number_of_nodes) * uint64(i))
	}

	// When the client ask the node
	key := "aasdfsasdaasdf132sdf313"
	hash := utils.GetHashValue(key)
	fmt.Println(hash)
	var idx int
	for i := number_of_nodes - 1; i >= 0; i-- {
		// 64* bit / number of nodes * (node id)
		if hash > node_value_mapping[i] {
			idx = i
			break
		}
	}
	// Ask if server idx is alive
	// Return idx value to the client
	fmt.Printf("\033[31mKey %v is stored in Node: %v\033[0m\n", key, idx)

	// create the router
	router := gin.Default()

	// mount router to routes
	routers.NodeHealth(router)

	router.Run(":" + os.Getenv("CENTRAL_SERVER_PORT"))
}
