package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	"calendly/config"
	"calendly/db"
	"calendly/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()

	gin.SetMode(gin.DebugMode)
}
