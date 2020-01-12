package main

import (
	"classifieds-api/internal/app/server"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	params := server.Params{}
	flag.Int64Var(&params.Port, "port", 5000, "http api port")
	flag.StringVar(&params.Url, "url", "/api", "url")
	flag.Parse()

	params.DatabaseURL = "host=localhost dbname=avito sslmode=disable port=5432 password=avito user=avito"

	err := server.StartApp(params)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
