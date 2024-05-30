package main

import (
	"learn-gin/router"
)

func main() {
	r := router.SetupRouter()
	_ = r.Run()
}
