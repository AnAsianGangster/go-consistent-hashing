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
	"bytes"
	"encoding/json"
	"fmt"
	"go-consistent-hashing/hintedHandoff"
	"go-consistent-hashing/nodeStatus"
	"go-consistent-hashing/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func FindOneKeyValuePair() gin.HandlerFunc {
	return func(context *gin.Context) {
		key := context.Query("key")
		// constructing the query URL
		var numOfAliveNodes = nodeStatus.GetNumberOfAliveNodes()
		var nodeLocation = utils.GetNodeLocation(numOfAliveNodes, key)
		var nodeName = nodeStatus.NodeIdxNameMap[nodeLocation]
		fmt.Printf(utils.ANSI_YELLOW+"%s"+utils.ANSI_RESET+"\n", nodeName)
		var port = nodeStatus.NodesStatus[nodeName].Port

		resp, err := http.Get("http://"+nodeName+":"+port+"/key-value-pair?key="+key)
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

		log.Println(string(body))

		context.JSON(200, gin.H{
			"success":  "find one endpoint hit",
			"location": nodeLocation,
			"data": string(body),
		})
	}
}

func CreatOneKeyValuePair() gin.HandlerFunc {
	return func(context *gin.Context) {
		// get key value pair from request body
		var keyValuePair KeyValuePair
		err := context.BindJSON(&keyValuePair)
		if err != nil {
			log.Fatal(err)
		}
		// constructing the query URL
		var numOfAliveNodes = nodeStatus.GetNumberOfAliveNodes()
		// primary node location
		var nodeLocation = utils.GetNodeLocation(numOfAliveNodes, keyValuePair.Key)
		// todo - better structure -> hardcoded the mod algo
		// todo - better structure -> hardcoded the replication -> put at meta data
		// nodeLocation is int
		storageNodes := make([]int, 0)
		responseBodies := make([]string, 0)
		for i := 0; i < 3; i++ {
			nodeNaiveLocation := nodeLocation + i
			curNodeLocation := nodeNaiveLocation % numOfAliveNodes
			storageNodes = append(storageNodes, curNodeLocation)

			// send the data to the nodes
			var curNodeName = nodeStatus.NodeIdxNameMap[curNodeLocation]
			var curPort = nodeStatus.NodesStatus[curNodeName].Port
			// request body
			requestBody, err := json.Marshal(map[string]string{
				"node": curNodeName,
				"key": keyValuePair.Key,
				"value": keyValuePair.Value,
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

			responseBodies = append(responseBodies, string(body))
		}

		context.JSON(200, gin.H{
			"success": "create one endpoint hit",
			"primary location": nodeLocation,
			"all nodes": fmt.Sprint(storageNodes),
			"data": fmt.Sprint(responseBodies),
		})

		// ========== hinted handoff ============
		if nodeStatus.GetNumberOfAliveNodes() != hintedHandoff.HARDCODED_NUMBER_OF_NODES {
			var hintedHandoffNodeLocation = utils.GetNodeLocation(hintedHandoff.HARDCODED_NUMBER_OF_NODES, keyValuePair.Key)
			hintedHandoff.CachedData[hintedHandoffNodeLocation] = append(hintedHandoff.CachedData[hintedHandoffNodeLocation], hintedHandoff.KeyValuePair{
				Key:   keyValuePair.Key,
				Value: keyValuePair.Value,
			})
			fmt.Println(hintedHandoff.CachedData)
		}
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
