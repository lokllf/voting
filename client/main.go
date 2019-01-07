package main

import (
	"voting/client/router"
)

func main() {
	r := router.Init()
	r.Run(":9009")
}
