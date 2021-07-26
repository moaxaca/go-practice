package controllers

import (
	"github.com/gin-gonic/gin"
)

// HealthCheckResponseSuccess represents the user for this application
//
// A user is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A user can have friends with whom they can share what they like.
//
// swagger:response health_check_response_success
type HealthCheckResponseSuccess struct {
	Message string `json:"message,omitempty"`
	Version string `json:"version,omitempty"`
	BuildSha string `json:"sha,omitempty"`
}

type infrastructureHandlers struct{}

// swagger:route GET /healthz Infrastructure healthz
//
// HealthCheck
//
// Health checks the application to ensure everything is setup and
// running correctly.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       default: server_error
//       200: health_check_response_success
//       401: access_error
//       404: miss_resource_error
//       422: validation_error
func (a *infrastructureHandlers) healthCheck(c *gin.Context) {
	response := HealthCheckResponseSuccess{}
	response.BuildSha = "na"
	response.Message = "App is healthy"
	response.Version = "0.0.0"
	c.JSON(200, response)
}

func RegisterInfrastructureRoutes(router *gin.RouterGroup) {
	handler := infrastructureHandlers{}
	router.GET("/healthz", handler.healthCheck)
}
