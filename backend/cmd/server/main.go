/*
 * @Author: SpenserCai
 * @Date: 2025-06-23 16:12:47
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2025-06-23 16:15:02
 * @Description: file content
 */
package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"

	"github.com/SpenserCai/OnceHumanTools/backend/restapi"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewOncehumanToolsAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "OnceHuman工具集API"
	parser.LongDescription = "提供OnceHuman游戏相关的计算工具API"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		log.Fatalf("Error parsing flags: %v (code: %d)", err, code)
	}

	server.ConfigureAPI()

	// 启动前打印信息
	log.Printf("Starting OnceHuman Tools API Server...")
	log.Printf("API Version: 1.0.0")
	log.Printf("Listening on %s:%d", server.Host, server.Port)
	log.Printf("Swagger UI available at: http://%s:%d/docs", server.Host, server.Port)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
