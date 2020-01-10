package main

import (
	"classifieds-api/internal/app/server"
	"flag"
	"fmt"
)

func main() {
	params := server.Params{}
	flag.Int64Var(&params.Port, "port", 5000, "http api port")
	flag.StringVar(&params.Url, "url", "/api", "url")
	flag.Parse()

	fmt.Println(params)
}
