package core

import (
	"ping/initialize"
)

func RunServer() {
	Router := initialize.Routers()
	Router.Run(":8080")
}