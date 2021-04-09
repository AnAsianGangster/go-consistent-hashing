/*
 * Author: Yang Aobo
 * Telegram: @AnAsianGangster
 * Created At: Apr 9, 2021
 * Updated At: Apr 9, 2021
 * Last Modified By: Yang Aobo
 */

/**
 * This package contains functions mount HTTP handlers & middlewares to URL routes
 *
 *
 * This file contains a function mount handlers to "/key-value-pair" route
 */
package routers

import (
	"go-consistent-hashing/controllers"

	"github.com/gin-gonic/gin"
)

// RESTful
func MountKeyValuePairIORouter(router *gin.Engine) *gin.Engine {
	router.GET("/key-value-pair", controllers.FindOneKeyValuePair())
	router.POST("/key-value-pair", controllers.CreatOneKeyValuePair())
	router.PUT("/key-value-pair", controllers.UpdateOneKeyValuePair())
	router.DELETE("/key-value-pair", controllers.DeleteOneKeyValuePair())
	return router
}
