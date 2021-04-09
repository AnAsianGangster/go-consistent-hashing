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
	"log"

	"github.com/gin-gonic/gin"
)

// TODO change to enums & move to a private folder
var NodesStatus = make(map[string]NodeStatusStruct)

// TODO I use Content-Type: application/json. Might need to change to postform
type NodeStatusStruct struct {
	Name     string `json:"name"`
	NodeName string `json:"nodeName"`
	Port     string `json:"port"`
	Status   string `json:"status"`
}

func ChangeNodeStatus() gin.HandlerFunc {
	return func(context *gin.Context) {
		var NodeStatusStruct NodeStatusStruct
		err := context.BindJSON(&NodeStatusStruct)
		if err != nil {
			log.Fatal(err)
		}

		NodesStatus[NodeStatusStruct.Name] = NodeStatusStruct

		fmt.Printf("\033[31m")
		fmt.Println(NodesStatus)
		fmt.Println("\033[0m")

		context.JSON(200, gin.H{
			"success": true,
		})
	}
}
