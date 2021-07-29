// Package internal Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v2
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//     oauth2:
//         type: oauth2
//         authorizationUrl: /oauth2/auth
//         tokenUrl: /oauth2/token
//         in: header
//         scopes:
//           bar: foo
//         flow: accessCode
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package internal

import (
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	"io.parcely.address_validation/internal/controllers"
	"io.parcely.address_validation/pkg/address_validator"
	"log"
	"os"
)

type RestServerConfiguration struct {
	Name string
	Address string
}

func CreateRestServer(config RestServerConfiguration) server.Server {
	// Create Micro Server
	srv := httpServer.NewServer(
		server.Name(config.Name),
		server.Address(config.Address),
	)

	// Create Router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// Register Middleware

	// Register Routes
	smartyCredentials := address_validator.SmartyStreetsCredentials{}
	smartyCredentials.AuthId = os.Getenv("SMART_AUTH_ID")
	smartyCredentials.AuthToken = os.Getenv("SMART_AUTH_TOKEN")
	addressValidatorBuilder := address_validator.CreateBuilder()
	addressValidatorBuilder.WithInMemoryCache()
	addressValidatorBuilder.WithSmartyValidator(smartyCredentials)
	addressValidator, _ := addressValidatorBuilder.Build()
	controllers.RegisterAddressRoutes(router.Group("/validate"), addressValidator)
	controllers.RegisterExampleRoutes(router.Group("/example"))
	controllers.RegisterInfrastructureRoutes(router.Group("/"))

	// Register Router Handler
	if err := srv.Handle(srv.NewHandler(router)); err != nil {
		log.Fatalln(err)
	}

	return srv
}