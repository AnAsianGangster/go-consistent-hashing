/*
 * Author: Yang Aobo
 * Telegram: @AnAsianGangster
 * Created At: Apr 9, 2021
 * Updated At: Apr 9, 2021
 * Last Modified By: Yang Aobo
 */

/**
 * This package contains HTTP handler functions
 *
 *
 * This file contains handler functions that handle I/O operations on key value pairs
 *
 * All functions destructure HTTP requests, forward the request to the node
 */
package controllers

import (
	"fmt"
	"go-consistent-hashing/nodeStatus"
	"go-consistent-hashing/utils"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// TODO forward the traffic to the node
func FindOneKeyValuePair() gin.HandlerFunc {
	return func(context *gin.Context) {
		key := context.Query("key")
		// constructing the query URL
		var numOfAliveNodes int = nodeStatus.GetNumberOfAliveNodes()
		var nodeLocation int = utils.GetNodeLocation(numOfAliveNodes, key)
		var nodeName string = nodeStatus.NodeIdxNameMap[nodeLocation]
		fmt.Printf(utils.ANSI_YELLOW+"%s"+utils.ANSI_RESET+"\n", nodeName)
		var port string = nodeStatus.NodesStatus[nodeName].Port

		resp, err := http.Get("http://"+nodeName+":"+port+"/key-value-pair?key="+key)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(body))

		context.JSON(200, gin.H{
			"success":  "find one endpoint hit",
			"location": nodeLocation,
		})
	}
}

func CreatOneKeyValuePair() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"success": "create one endpoint hit",
		})
	}
}

func UpdateOneKeyValuePair() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"success": "update one endpoint hit",
		})
	}
}

func DeleteOneKeyValuePair() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"success": "delete one endpoint hit",
		})
	}
}
