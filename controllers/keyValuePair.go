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
	"github.com/gin-gonic/gin"
)

// TODO forward the traffic to the node
func FindOneKeyValuePair() gin.HandlerFunc {
		return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"success": "find one endpoint hit",
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
