/*
 * Author: Yang Aobo
 * Telegram: @AnAsianGangster
 * Created At: March 29, 2021
 * Updated At: March 29, 2021
 * Last Modified By: Yang Aobo
 */

/**
 * This package contains functions mount HTTP handlers & middlewares to URL routes
 *
 *
 * This file contains a function mount handlers to "/node-health" route
 */
package routers

import (
	"go-consistent-hashing/controllers"

	"github.com/gin-gonic/gin"
)

// RESTful
func NodeHealth(router *gin.Engine) *gin.Engine {
	router.POST("/node-status", controllers.ChangeNodeStatus())
	return router
}
