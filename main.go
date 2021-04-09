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
	// ping for nodes concurrently
	go func() {
		for {
			nodeStatus.UpdateNodesStatusMap()
			// fmt.Println(time.Now().UTC())
			// sleep for one second
			time.Sleep(time.Second * 1)
		}
	}()
	
	// TODO Listen at some port
	// TODO mapping for node name with index in the node_value_mapping
	number_of_nodes := len(controllers.NodesStatus)
	fmt.Println(utils.GetNodeLocation(number_of_nodes, "qwertyuio"))

	// create the router
	router := gin.Default()

	// mount router to routes
	routers.NodeHealth(router)
	routers.MountKeyValuePairIORouter(router)

	router.Run(":" + os.Getenv("CENTRAL_SERVER_PORT"))
}
