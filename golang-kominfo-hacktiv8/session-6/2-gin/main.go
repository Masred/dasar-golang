package main

import "github.com/masred/dasar-golang/golang-kominfo-hactiv8/session-6/2-gin/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
