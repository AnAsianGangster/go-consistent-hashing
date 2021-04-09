package main

import (
	"fmt"
	"go-consistent-hashing/nodeStatus"
	"go-consistent-hashing/routers"
	"go-consistent-hashing/utils"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// ping for nodes concurrently
	go func() {
		for {
			nodeStatus.UpdateNodesStatusMap()
			// sleep for one second
			time.Sleep(time.Second * 1)
		}
	}()

	// TODO Listen at some port
	number_of_nodes := nodeStatus.GetNumberOfAliveNodes()
	var nodeIdxNameMap = make(map[int]string)
	idx := 0
	for _, val := range nodeStatus.NodesStatus {
		nodeIdxNameMap[idx] = val.NodeName
		idx++
	}
	fmt.Println(nodeIdxNameMap)
	fmt.Println(utils.GetNodeLocation(number_of_nodes, "qwertyuio"))

	// create the router
	router := gin.Default()

	// mount router to routes
	routers.NodeHealth(router)
	routers.MountKeyValuePairIORouter(router)

	router.Run(":" + os.Getenv("CENTRAL_SERVER_PORT"))
}
