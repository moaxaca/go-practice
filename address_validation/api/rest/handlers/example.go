package handlers

import (
	"github.com/gin-gonic/gin"
)

// ExampleResponseSuccess represents the user for this application
//
// A user is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A user can have friends with whom they can share what they like.
//
// swagger:response example_response_success
type ExampleResponseSuccess struct {
	Message string `json:"name,omitempty"`
}

type exampleHandlers struct{}

// swagger:route GET /example Example exampleGet
//
// Example Get Request
//
// This is a working example of how to create an endpoint.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//       api_key:
//
//     Responses:
//       default: server_error
//       200: example_response_success
//       401: access_error
func (a *exampleHandlers) index(c *gin.Context) {
	response := ExampleResponseSuccess{Message: "Example Endpoint"}
	c.JSON(200, response)
}

func RegisterExampleRoutes(router *gin.RouterGroup) {
	handler := exampleHandlers{}
	router.GET("/", handler.index)
}
