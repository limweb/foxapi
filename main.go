package main

import (
	"apitest/db"
	"apitest/webserver"
	"flag"
	"log"
)

var (
	port string
)

func init() {
	log.Println("----init start-----")
	flag.StringVar(&port, "port", "8080", "web port 8080")
}

// @title ApiTest
// @version 1.0
// @description.markdown
// @termsOfService http://www.servit.co.th/

// @contact.name API Support
// @contact.url http://www.servit.co.th/support
// @contact.email limweb@hotmail.com

// @license.name The MIT License
// @license.url https://opensource.org/licenses/MIT

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flag.Parse()
	db.SetupDB()
	port = ":" + port
	log.Printf("----------------------  Web Server Api  V 0.0.1   --------------------------")
	log.Printf("----------------------วิธีใช้: server.exe -port 8080 --------------------------")
	log.Printf("-------------------Server is starting at 127.0.0.1:%s -------------------")
	webserver.StartWeb(port)
}
