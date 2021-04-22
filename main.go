package main

import (
	"fmt"
	"go-consistent-hashing/nodeStatus"
	"go-consistent-hashing/routers"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// ping for nodes concurrently
	go func() {
		for {
			nodeStatus.UpdateNodesStatusMap()
			// FIXME this mapping changes
			nodeStatus.UpdateNodesStatusMapToArrayMapping()
			// sleep for one second
			time.Sleep(time.Second * 1)
		}
	}()

	// create the router
	router := gin.Default()

	// mount router to routes
	routers.NodeHealth(router)
	routers.MountKeyValuePairIORouter(router)

	router.Run(":" + os.Getenv("CENTRAL_SERVER_PORT"))

}

func checkNodes() {
	uptimeTicker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-uptimeTicker.C:
			ports := [5]string{"5000", "5001", "5002", "5003", "5004"}
			for _, port := range ports {
				if nodeStatus.GetOneNodeStatus(port) != "AlIVE" {
					log.Println(fmt.Sprintf("Node %s is dead", port))
				}
			}
		}
	}
}