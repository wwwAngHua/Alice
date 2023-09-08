package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route = CollectRoute(route)
	panic(route.Run(":8080"))
}
