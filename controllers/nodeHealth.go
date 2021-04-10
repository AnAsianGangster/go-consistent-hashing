/*
 * Author: Yang Aobo
 * Telegram: @AnAsianGangster
 * Created At: April 7, 2021
 * Updated At: April 7, 2021
 * Last Modified By: Yang Aobo
 */

/**
 * This package contains HTTP handler functions
 *
 *
 * This file contains handler functions that handle node health related operations
 *
 * All functions destructure HTTP requests, call database operations, build response
 * and reply with response
 */
package controllers

import (
	"fmt"
	"go-consistent-hashing/nodeStatus"
	"log"

	"github.com/gin-gonic/gin"
)

func ChangeNodeStatus() gin.HandlerFunc {
	return func(context *gin.Context) {
		var NodeStatusStruct nodeStatus.NodeStatusStruct
		err := context.BindJSON(&NodeStatusStruct)
		if err != nil {
			log.Fatal(err)
		}

		nodeStatus.NodesStatus[NodeStatusStruct.NodeName] = NodeStatusStruct

		fmt.Printf("\033[31m")
		fmt.Println(nodeStatus.NodesStatus)
		fmt.Println("\033[0m")

		context.JSON(200, gin.H{
			"success": true,
		})
	}
}
