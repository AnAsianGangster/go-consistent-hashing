package main

import (
	"go-consistent-hashing/nodeStatus"
	"go-consistent-hashing/routers"
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
