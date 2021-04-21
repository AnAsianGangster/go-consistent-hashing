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
	"bytes"
	"encoding/json"
	"fmt"
	"go-consistent-hashing/hintedHandoff"
	"go-consistent-hashing/nodeStatus"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

		// ========== hinted handoff ============
		charNodeLocation := NodeStatusStruct.NodeName[len(NodeStatusStruct.NodeName) - 1:]
		intNodeLocation, err := strconv.Atoi(charNodeLocation)
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		// assume there is cached data is inside the map
		// todo - encapsulate below, now it's hardcoded here
		dataHandedOff := make([]string, 0)
		for _, curCachedData := range hintedHandoff.CachedData[intNodeLocation] {
			var curNodeName = nodeStatus.NodeIdxNameMap[intNodeLocation]
			var curPort = nodeStatus.NodesStatus[curNodeName].Port
			// request body
			requestBody, err := json.Marshal(map[string]string{
				"node": curNodeName,
				"key": curCachedData.Key,
				"value": curCachedData.Value,
			})

			resp, err := http.Post("http://"+curNodeName+":"+curPort+"/key-value-pair?", "application/json", bytes.NewBuffer(requestBody))
			if err != nil {
				log.Fatal(err)
			}

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(resp.Body)

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			dataHandedOff = append(dataHandedOff, string(body))
		}
		// print handed off data in the console
		fmt.Println(" _     _       _           _    _                     _        __  __")
		fmt.Println("| |__ (_)_ __ | |_ ___  __| |  | |__   __ _ _ __   __| | ___  / _|/ _|")
		fmt.Println("| '_ \\| | '_ \\| __/ _ \\/ _` |  | '_ \\ / _` | '_ \\ / _` |/ _ \\| |_| |_ ")
		fmt.Println("| | | | | | | | ||  __/ (_| |  | | | | (_| | | | | (_| | (_) |  _|  _|")
		fmt.Println("|_| |_|_|_| |_|\\__\\___|\\__,_|  |_| |_|\\__,_|_| |_|\\__,_|\\___/|_| |_|")
		fmt.Print("\u001B[35m") // purple
		fmt.Print(dataHandedOff)
		fmt.Print("\u001B[0m\n")
	}
}
